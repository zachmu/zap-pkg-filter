// Copyright 2024 Dolthub, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"os"
	"testing"
	"zap-pkg-filter/example/a"
	"zap-pkg-filter/example/a/c"
	"zap-pkg-filter/example/a/c/e"
	"zap-pkg-filter/example/b"
	"zap-pkg-filter/example/b/d"
	"zap-pkg-filter/pkglog"
)

func BenchmarkLogging(t *testing.B) {
	reset()
	for i := 0; i < t.N; i++ {
		main()
	}
}

func BenchmarkSparseLogging(t *testing.B) {
	reset()
	pkglog.InitFromFile("sparse_logging_props.txt")
	for i := 0; i < t.N; i++ {
		main()
	}
}

func reset() {
	os.Remove("output.log")
	a.ResetLogger()
	b.ResetLogger()
	c.ResetLogger()
	d.ResetLogger()
	e.ResetLogger()
}
