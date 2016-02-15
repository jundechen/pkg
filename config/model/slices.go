// Copyright 2015-2016, Cyrill @ Schumacher.fm and the CoreStore contributors
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

package model

import (
	"strconv"
	"strings"

	"github.com/corestoreio/csfw/config"
	"github.com/corestoreio/csfw/store/scope"
	"github.com/corestoreio/csfw/util"
	"github.com/corestoreio/csfw/util/bufferpool"
	"github.com/juju/errors"
)

// CSVSeparator separates CSV values. Default value.
const CSVSeparator = ","

// StringCSV represents a path in config.Getter which will be saved as a
// CSV string and returned as a string slice. Separator is a comma.
type StringCSV struct {
	// CSVSeparator is your custom separator. Defaults to CSVSeparator
	CSVSeparator string
	baseValue
}

// NewStringCSV creates a new CSV string type. Acts as a multiselect.
// Default separator: constant CSVSeparator
func NewStringCSV(path string, opts ...Option) StringCSV {
	return StringCSV{
		CSVSeparator: CSVSeparator,
		baseValue:    NewValue(path, opts...),
	}
}

// Get returns a string slice. Splits the stored string by comma.
// Can return nil,nil. Empty values will be discarded. Returns a slice
// containing unique entries. No validation will be made.
func (str StringCSV) Get(sg config.ScopedGetter) ([]string, error) {
	s, err := str.lookupString(sg)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}
	var ret util.StringSlice = strings.Split(s, str.CSVSeparator)
	return ret.Unique(), nil
}

// Write writes a slice with its scope and ID to the writer.
// Validates the input string slice for correct values if set in source.Slice.
func (str StringCSV) Write(w config.Writer, sl []string, s scope.Scope, scopeID int64) error {
	for _, v := range sl {
		if err := str.ValidateString(v); err != nil {
			return err
		}
	}
	return str.baseValue.Write(w, strings.Join(sl, str.CSVSeparator), s, scopeID)
}

// IntCSV represents a path in config.Getter which will be saved as a
// CSV string and returned as an int64 slice. Separator is a comma.
type IntCSV struct {
	baseValue
	// Lenient ignores errors in parsing integers
	Lenient bool
	// CSVSeparator custom separator, default is constant CSVSeparator
	CSVSeparator string
}

// NewIntCSV creates a new int CSV type. Acts as a multiselect.
func NewIntCSV(path string, opts ...Option) IntCSV {
	return IntCSV{
		baseValue:    NewValue(path, opts...),
		CSVSeparator: CSVSeparator,
	}
}

// Get returns an int slice. Int string gets splited by comma.
// Can return nil,nil. If multiple values cannot be casted to int then the
// last known error gets returned.
func (ic IntCSV) Get(sg config.ScopedGetter) ([]int, error) {
	s, err := ic.lookupString(sg)
	if err != nil {
		return nil, err
	}
	if s == "" {
		return nil, nil
	}

	csv := strings.Split(s, ic.CSVSeparator)
	ret := make([]int, 0, len(csv))

	for _, line := range csv {
		v, err := strconv.Atoi(line)
		if err != nil && false == ic.Lenient {
			return ret, err
		}
		if err == nil {
			ret = append(ret, v)
		}
	}
	return ret, nil
}

// Write writes int values as a CSV string
func (ic IntCSV) Write(w config.Writer, sl []int, s scope.Scope, scopeID int64) error {

	val := bufferpool.Get()
	defer bufferpool.Put(val)
	for i, v := range sl {

		if err := ic.ValidateInt(v); err != nil {
			return err
		}

		if _, err := val.WriteString(strconv.Itoa(v)); err != nil {
			return errors.Mask(err)
		}
		if i < len(sl)-1 {
			if _, err := val.WriteString(ic.CSVSeparator); err != nil {
				return errors.Mask(err)
			}
		}
	}
	return ic.baseValue.Write(w, val.String(), s, scopeID)
}
