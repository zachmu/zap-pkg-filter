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

package pkglog

import (
	"os"

	"go.uber.org/zap"
	"moul.io/zapfilter"
)

var filterRules zapfilter.FilterFunc

func Init(config string) {
	filterRules = zapfilter.MustParseRules(config)
}

func InitFromFile(configFile string) {
	file, err := os.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	filterRules = zapfilter.MustParseRules(string(file))
}

func NewProductionLogger(namespace string) *zap.Logger {
	config := zap.NewProductionConfig()
	config.OutputPaths = []string{"output.log"}
	config.Sampling = nil
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	underlying, err := config.Build()
	if err != nil {
		panic(err)
	}

	return zap.New(zapfilter.NewFilteringCore(underlying.Core(), filterRules)).Named(namespace)
}

func NewDevelopmentLogger(namespace string) *zap.Logger {
	underlying, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	return zap.New(zapfilter.NewFilteringCore(underlying.Core(), filterRules)).Named(namespace)
}
