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
	"strings"

	"github.com/corestoreio/pkg/util/conv"
	"github.com/corestoreio/pkg/util/strs"
)

type common struct {
	packageName string
	*bytes.Buffer
	packageNames map[string]string // Imported package names in the current file.
	indent       string
}

// AddImport adds a new import path. importPath required and packageName optional.
func (g *common) AddImport(importPath, packageName string) {
	g.packageNames[importPath] = packageName
}

// Writes a multiline comment and formats it to a max width of 80 chars. It adds
// automatically the comment prefix `//`.
func (g *common) C(comments ...string) {
	cLines := strings.Split(strs.WordWrap(strings.Join(comments, " "), 78), "\n")
	for _, c := range cLines {
		g.WriteString("// ")
		g.WriteString(c)
		g.WriteByte('\n')
	}
}

// P prints the arguments to the generated output. It tries to convert all kind
// of types to a string.
func (g *common) P(str ...interface{}) {
	_, _ = g.WriteString(g.indent)
	for _, v := range str {
		s, err := conv.ToStringE(v)
		if err != nil {
			panic(err)
		}
		_, _ = g.WriteString(s)
		g.WriteByte(' ')
	}
	_ = g.WriteByte('\n')
}

// In Indents the output one tab stop.
func (g *common) In() { g.indent += "\t" }

// Out unindents the output one tab stop.
func (g *common) Out() {
	if len(g.indent) > 0 {
		g.indent = g.indent[1:]
	}
}
