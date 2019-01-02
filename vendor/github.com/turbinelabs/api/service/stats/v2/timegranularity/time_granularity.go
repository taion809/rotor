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

// This file was automatically generated by doc.go from github.com/turbinelabs/api/service/stats/enum.template.
// Any changes will be lost if this file is regenerated.

package timegranularity

import (
	"encoding/json"
	"fmt"
)

// TimeGranularity is an enumeration.
type TimeGranularity int

// Defined values of TimeGranularity.
const (
	Minutes TimeGranularity = iota
	Hours
	Unknown
)

var _dummy = TimeGranularity(0)
var _ json.Marshaler = &_dummy
var _ json.Unmarshaler = &_dummy

const (
	strMinutes = "minutes"
	strHours   = "hours"
	strUnknown = "unknown"
)

var timeGranularityNames = [...]string{
	strMinutes,
	strHours,
	strUnknown,
}

const minTimeGranularity = TimeGranularity(0)

var maxTimeGranularity = TimeGranularity(len(timeGranularityNames) - 2)

// IsValid returns a boolean indicating whether the given TimeGranularity is defined
// (valid) or not.
func IsValid(i TimeGranularity) bool {
	return i >= minTimeGranularity && i <= maxTimeGranularity
}

// FromName converts a TimeGranularity name into a TimeGranularity.
// Returns Unknown if the name is not known. Names are case sensitive.
func FromName(s string) TimeGranularity {
	for idx, name := range timeGranularityNames {
		candidate := TimeGranularity(idx)
		if IsValid(candidate) && name == s {
			return candidate
		}
	}

	return Unknown
}

// ForEach invokes the given function for each valid TimeGranularity.
func ForEach(f func(TimeGranularity)) {
	for i := int(minTimeGranularity); i <= int(maxTimeGranularity); i++ {
		tg := TimeGranularity(i)
		f(tg)
	}
}

// String return this TimeGranularity's string representation.
func (i TimeGranularity) String() string {
	if !IsValid(i) {
		return fmt.Sprintf("unknown(%d)", i)
	}
	return timeGranularityNames[i]
}

// MarshalJSON converts this TimeGranularity to a quoted JSON string. Returns an
// error if the TimeGranularity is nil or invalid.
func (i *TimeGranularity) MarshalJSON() ([]byte, error) {
	if i == nil {
		return nil, fmt.Errorf("cannot marshal unknown TimeGranularity (nil)")
	}

	qt := *i
	if !IsValid(qt) {
		return nil, fmt.Errorf("cannot marshal unknown TimeGranularity (%d)", qt)
	}

	name := timeGranularityNames[qt]
	b := make([]byte, 0, len(name)+2)
	b = append(b, '"')
	b = append(b, name...)
	return append(b, '"'), nil
}

// UnmarshalJSON converts a quoted JSON string into a TimeGranularity. Returns an
// error if the receiver is nil, the JSON is not a quoted string, or if the
// string does not represent a valid TimeGranularity. Otherwise, the receiver's value
// is set to the TimeGranularity represented by the string.
func (i *TimeGranularity) UnmarshalJSON(bytes []byte) error {
	if i == nil {
		return fmt.Errorf("cannot unmarshal into nil TimeGranularity")
	}

	length := len(bytes)
	if length <= 2 || bytes[0] != '"' || bytes[length-1] != '"' {
		return fmt.Errorf("cannot unmarshal invalid JSON: %q", string(bytes))
	}

	unmarshalName := string(bytes[1 : length-1])

	qt := FromName(unmarshalName)
	if qt == Unknown {
		return fmt.Errorf(
			"cannot unmarshal unknown TimeGranularity %q",
			unmarshalName,
		)
	}

	*i = qt
	return nil
}

// UnmarshalForm converts a string into a TimeGranularity. Returns an error if the
// receiver is nil or if the string does not represent a valid TimeGranularity.
// Otherwise, the receiver's value is set to the TimeGranularity represented by the
// string.
func (i *TimeGranularity) UnmarshalForm(value string) error {
	if i == nil {
		return fmt.Errorf("cannot unmarshal into nil TimeGranularity")
	}

	qt := FromName(value)
	if qt == Unknown {
		return fmt.Errorf("cannot unmarshal unknown TimeGranularity %q", value)
	}

	*i = qt
	return nil
}