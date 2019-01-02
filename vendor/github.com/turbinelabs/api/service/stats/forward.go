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

package stats

import (
	v2 "github.com/turbinelabs/api/service/stats/v2"
)

// Payload is an alias for the V2 stats API Payload type.
type Payload = v2.Payload

// Stat is an alias for the V2 stats API Stat type.
type Stat = v2.Stat

// Histogram is an alias for the V2 stats API Histogram type.
type Histogram = v2.Histogram

// ForwardResult is an alias for the V2 stats API ForwardResult type
type ForwardResult = v2.ForwardResult
