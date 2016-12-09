// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.
//
// Author: Andrei Matei (andreimatei1@gmail.com)

// Functions used for testing metrics.

package sql_test

import (
	"github.com/cockroachdb/cockroach/pkg/testutils/serverutils"
	"github.com/cockroachdb/cockroach/pkg/util/metric"
	"github.com/pkg/errors"
)

func checkCounterEQ(s serverutils.TestServerInterface, meta metric.Metadata, e int64) error {
	if a := s.MustGetSQLCounter(meta.Name); a != e {
		return errors.Errorf("stat %s: actual %d != expected %d", meta.Name, a, e)
	}
	return nil
}

func checkCounterGE(s serverutils.TestServerInterface, meta metric.Metadata, e int64) error {
	if a := s.MustGetSQLCounter(meta.Name); a < e {
		return errors.Errorf("stat %s: expected: actual %d >= %d", meta.Name, a, e)
	}
	return nil
}
