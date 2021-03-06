/*
Copyright 2018 Turbine Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/turbinelabs/api"
	apihttp "github.com/turbinelabs/api/http"
	httperr "github.com/turbinelabs/api/http/error"
	"github.com/turbinelabs/api/queryargs"
	"github.com/turbinelabs/api/service/changelog"
	tbntime "github.com/turbinelabs/nonstdlib/time"
)

type httpHistoryV1 struct {
	dest apihttp.Endpoint

	requestHandler apihttp.RequestHandler
}

// Construct a new HTTP backed History API implementation.
//
// Parameters:
//	dest - service handling our HTTP requests; cf. NewService
func NewHistoryV1(dest apihttp.Endpoint) (*httpHistoryV1, error) {
	return &httpHistoryV1{
		dest,
		apihttp.NewRequestHandler(dest.Client()),
	}, nil
}

// creates a cluster-scoped version of the specified path
func (hh *httpHistoryV1) path(p string) string {
	return "/v1.0/changelog" + p
}

// Construct a request to the associated cluster Endpoint with a specified
// method, path, query params, and body.
func (hh *httpHistoryV1) request(
	method string,
	path string,
	params apihttp.Params,
	body string,
) (*http.Request, error) {
	rdr := strings.NewReader(body)
	req, err := hh.dest.NewRequest(string(method), hh.path(path), params, rdr)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func (hh *httpHistoryV1) get(path string, params apihttp.Params) (*http.Request, error) {
	return hh.request(http.MethodGet, path, params, "")
}

func (hh *httpHistoryV1) Index(
	filters changelog.FilterExpr,
	start,
	end time.Time,
) ([]api.ChangeDescription, error) {
	b, err := json.Marshal(filters)
	if err != nil {
		return nil, httperr.New400(
			fmt.Sprintf("unable to encode user filters: %v", filters),
			httperr.UnknownUnclassifiedCode)
	}

	params := apihttp.Params{
		queryargs.IndexFilters: string(b),
	}

	if !start.IsZero() {
		params[queryargs.WindowStart] = fmt.Sprintf("%v", tbntime.ToUnixMicro(start))
	}
	if !end.IsZero() {
		params[queryargs.WindowStop] = fmt.Sprintf("%v", tbntime.ToUnixMicro(end))
	}

	var response []api.ChangeDescription
	reqFn := func() (*http.Request, error) { return hh.get("/adhoc", params) }
	if err := hh.requestHandler.Do(reqFn, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func mkTimeArgs(start, stop time.Time) apihttp.Params {
	result := apihttp.Params{}
	fillTime := func(n string, t time.Time) {
		if t.IsZero() {
			return
		}
		s := tbntime.ToUnixMicro(t)
		result[n] = fmt.Sprintf("%v", s)
	}

	fillTime(queryargs.WindowStart, start)
	fillTime(queryargs.WindowStop, stop)
	return result
}

func doNamedChangelogQueryRequest(
	hh *httpHistoryV1,
	url string,
	start,
	stop time.Time,
) ([]api.ChangeDescription, error) {
	params := mkTimeArgs(start, stop)
	var response []api.ChangeDescription

	reqFn := func() (*http.Request, error) { return hh.get(url, params) }
	if err := hh.requestHandler.Do(reqFn, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (hh *httpHistoryV1) DomainGraph(
	domainKey api.DomainKey,
	start,
	stop time.Time,
) ([]api.ChangeDescription, error) {
	url := fmt.Sprintf("/domain-graph/%s", string(domainKey))
	return doNamedChangelogQueryRequest(hh, url, start, stop)
}

func (hh *httpHistoryV1) RouteGraph(
	routeKey api.RouteKey,
	start,
	stop time.Time,
) ([]api.ChangeDescription, error) {
	url := fmt.Sprintf("/route-graph/%s", string(routeKey))
	return doNamedChangelogQueryRequest(hh, url, start, stop)
}

func (hh *httpHistoryV1) SharedRulesGraph(
	clusterKey api.SharedRulesKey,
	start,
	stop time.Time,
) ([]api.ChangeDescription, error) {
	url := fmt.Sprintf("/shared-rules-graph/%s", string(clusterKey))
	return doNamedChangelogQueryRequest(hh, url, start, stop)
}

func (hh *httpHistoryV1) ClusterGraph(
	clusterKey api.ClusterKey,
	start,
	stop time.Time,
) ([]api.ChangeDescription, error) {
	url := fmt.Sprintf("/cluster-graph/%s", string(clusterKey))
	return doNamedChangelogQueryRequest(hh, url, start, stop)
}

func (hh *httpHistoryV1) Zone(
	zoneKey api.ZoneKey,
	start,
	stop time.Time,
) ([]api.ChangeDescription, error) {
	url := fmt.Sprintf("/zone/%s", string(zoneKey))
	return doNamedChangelogQueryRequest(hh, url, start, stop)
}
