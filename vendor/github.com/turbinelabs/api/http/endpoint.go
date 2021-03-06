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

package http

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Protocol indicates which transport we should use for communicating with our service.
type Protocol string

const (
	// HTTP is the Protocol constant for HTTP requests
	HTTP Protocol = "http"

	// HTTPS is the Protocol constant for HTTP requests
	HTTPS Protocol = "https"
)

// Params holds URL query arg name -> value mappings
type Params map[string]string

// NewEndpoint constructs a new HTTP Endpoint. This is used to configure the
// HTTP service implementation.
//
// Parameters:
// 	protocol
//	  specifies whether we should be using HTTP or HTTPS there is currently no
// 	  special configuration for HTTPS (certificate pinning, custom root CAs,
//	  etc.)
//	hostPort - host:port for the service we should connect to
//
// Returns a new Endpoint object and an error if there was a problem. Currently
// the only error possible is the result of a failed call to url.Parse which
// will be passed directly to the caller.
//
// The Endpoint object is configured with no custom headers (see
// Endpoint.AddHeader), and the net/http.Client created by
// HeaderPreservingClient. You may specify an alternate client via
// Endpoint.SetClient.
func NewEndpoint(protocol Protocol, hostPort string) (Endpoint, error) {
	url, err := url.Parse(fmt.Sprintf("%s://%s", string(protocol), hostPort))
	if err != nil {
		return Endpoint{}, err
	}

	return Endpoint{
		hostPort: hostPort,
		protocol: protocol,
		header:   http.Header{},
		client:   HeaderPreservingClient(),
		urlBase:  url,
	}, nil
}

// Endpoint is a effectively a net/http.Request factory, which configures
// the common concerns of requests to a given HTTP service.
type Endpoint struct {
	hostPort string
	protocol Protocol
	header   http.Header
	client   *http.Client

	urlBase *url.URL // computed at construction
}

// Copy makes a copy of the Endpoint. Insures that modifications to custom
// headers of the new Endpoint are not made to the original Endpoint
// and vice versa.
func (e *Endpoint) Copy() Endpoint {
	headerCopy := make(http.Header, len(e.header))
	for header, values := range e.header {
		for _, value := range values {
			headerCopy.Add(header, value)
		}
	}

	newE := *e
	newE.header = headerCopy
	return newE
}

// Client returns the net/http.Client for this Endpoint.
func (e *Endpoint) Client() *http.Client {
	return e.client
}

// SetClient sets an alternative net/http.Client for this Endpoint.
func (e *Endpoint) SetClient(c *http.Client) {
	e.client = c
}

// AddHeader adds a header to be added to all requests created via NewRequest.
// These headers are meant to be constant across all requests (e.g. a client
// identifier). Headers specific to a particular request should be added
// directly to the net/http.Request.
func (e *Endpoint) AddHeader(header, value string) {
	e.header.Add(header, value)
}

// URL constructs a URL to this turbine Endpoint.
func (e *Endpoint) URL(path string, queryParams Params) string {
	newURL := *e.urlBase
	newURL.Path = path

	if len(queryParams) != 0 {
		q := newURL.Query()
		for k, v := range queryParams {
			q.Set(k, v)
		}
		newURL.RawQuery = q.Encode()
	}

	return newURL.String()
}

// NewRequest constructs a net/http.Request for this turbine Endpoint with the
// given method, path, (optional) query parameters and (optional) body. Headers
// previously configured via AddHeader are added automatically.
func (e *Endpoint) NewRequest(
	method string,
	path string,
	queryParams Params,
	body io.Reader,
) (*http.Request, error) {
	url := e.URL(path, queryParams)
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for header, values := range e.header {
		for _, value := range values {
			// header values are normalized when added
			if header == "Host" {
				request.Host = value
			}
			request.Header.Add(header, value)
		}
	}

	return request, nil
}
