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

package flag

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TimestampConversionFencepost represents an integer time value.
// Values larger than this are assumed to be in milliseconds since the
// Unix epoch. Treated as millisecons since the epoch, this timestamp
// represents 1973-01-11T16:26:24Z. Treated as seconds, this timestamp
// represents 5000-01-01:00:00:00Z.
const TimestampConversionFencepost = int64(95617584000)

// Timestamp conforms to the flag.Value and flag.Getter interfaces. It
// can be used to populate a timestamp from a command line
// argument. It accepts the following inputs:
//
// - timestamps in time.RFC3339Nano format (fractional seconds
//   optional)
//
// - integer seconds since the Unix epoch (see
//   TimestampConversionFencepost)
//
// - integer milliseconds since the Unix epoch (see
//   TimestampConversionFencepost)
//
// - "now" (case insensitive)
type Timestamp struct {
	Value time.Time
}

var _ flag.Getter = &Timestamp{}

// NewTimestamp creates a new Timestamp with the given default time.
func NewTimestamp(defaultTime time.Time) Timestamp {
	return Timestamp{Value: defaultTime}
}

// Set sets the current value of the Timestamp.
func (t *Timestamp) Set(value string) error {
	if strings.ToLower(value) == "now" {
		t.Value = time.Now()
		return nil
	} else if ticks, err := strconv.ParseInt(value, 10, 64); err == nil {
		if ticks >= TimestampConversionFencepost {
			t.Value = time.Unix(ticks/1000, (ticks%1000)*int64(time.Millisecond)).UTC()
		} else {
			t.Value = time.Unix(ticks, 0).UTC()
		}
		return nil
	} else if ts, err := time.Parse(time.RFC3339Nano, value); err == nil {
		t.Value = ts
		return nil
	}

	return fmt.Errorf(
		"cannot parse '%s': expecting seconds or milliseconds since the Unix epoch or RFC3339 format (fractional seconds optional)",
		value,
	)
}

// Get retrieves the current value of the Timestamp.
func (t *Timestamp) Get() interface{} {
	return t.Value
}

// String returns the current value of the Timestamp as a string in
// RFC33339 Nano format.
func (t *Timestamp) String() string {
	return t.Value.Format(time.RFC3339Nano)
}
