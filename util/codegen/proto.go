// Copyright 2015-present, Cyrill @ Schumacher.fm and the CoreStore contributors
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

package codegen

import (
	"bytes"
	"fmt"
	"io"
	"sort"

	"github.com/corestoreio/errors"
)

type Proto struct {
	common
	options []string
}

// NewProto creates a new Proto source code generator for a specific new package.
func NewProto(packageName string) *Proto {
	return &Proto{
		common: common{
			Buffer:       new(bytes.Buffer),
			packageName:  packageName,
			packageNames: map[string]string{},
		},
	}
}

func (g *Proto) AddOptions(nameRawValue ...string) {
	if len(nameRawValue)%2 == 1 {
		panic(errors.Fatal.Newf("[codegen] slice nameRawValue must be balanced: key, value"))
	}
	g.options = append(g.options, nameRawValue...)
}

func (g *Proto) generateImports(w io.Writer) {
	pkgSorted := make([]string, 0, len(g.packageNames))
	for key := range g.packageNames {
		pkgSorted = append(pkgSorted, key)
	}
	sort.Strings(pkgSorted)
	for _, p := range pkgSorted {
		fmt.Fprintf(w, "import %q;\n", p)
	}
}

func (g *Proto) generateOptions(w io.Writer) {
	sort.Strings(g.options)
	for i := 0; i < len(g.options); i += 2 {
		k := g.options[i]
		v := g.options[i+1]
		fmt.Fprintf(w, "option %s = %s;\n", k, v)
	}
}

func (g *Proto) GenerateFile(w io.Writer) error {

	var buf bytes.Buffer

	fmt.Fprintln(&buf, "// Auto generated source code")
	fmt.Fprintln(&buf, `syntax = "proto3";`)
	fmt.Fprintf(&buf, "package %s;\n", g.packageName)
	g.generateImports(&buf)
	g.generateOptions(&buf)

	g.Buffer.WriteTo(&buf)

	_, err := buf.WriteTo(w)
	return err
}
