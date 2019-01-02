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

package io

import (
	"errors"
	"io"
)

const FailingReaderMessage = "failingReader error"

type failingReader struct{}

// NewFailingReader produces an io.ReadCloser that returns an error for all
// calls to Read()
func NewFailingReader() io.ReadCloser {
	return &failingReader{}
}

func (r *failingReader) Read(p []byte) (int, error) {
	return 0, errors.New(FailingReaderMessage)
}

func (r *failingReader) Close() error {
	return nil
}
