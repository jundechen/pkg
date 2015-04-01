// Copyright 2015 CoreStore Authors
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

import "github.com/corestoreio/csfw/tools"

/* @todo
   Data will be "carved in stone" because it only changes during development.
   - attribute_set related tables: eav_attribute_set, eav_entity_attribute, eav_attribute_group, etc
   - label and option tables will not be hard coded
*/

const tplTypeDefinition = `
type (
    // @todo website must be present in the slice
    // {{.Name | prepareVar}} a data container for the data from a MySQL query
    {{.Name | prepareVar | toLowerFirst}} struct {
        {{ range .Columns }}{{.GoName | toLowerFirst}} {{.GoType}}
        {{ end }} }
)

{{ range .Columns }} func (a *{{$.Name | prepareVar | toLowerFirst}}) {{.GoName}}() {{.GoType}}{
    return a.{{.GoName | toLowerFirst}}
}
{{ end }}

// Check if Attributer interface has been successfully implemented
var _ {{.EAVPackage}}.Attributer = (*{{.Name | prepareVar | toLowerFirst}})(nil)

`

// here iota must start with 0 because constants are used as slice index.
const tplTypeDefinitionFile = tools.Copyright + `
package {{ .PackageName }}
    import (
        "github.com/corestoreio/csfw/eav"
        "github.com/corestoreio/csfw/{{ .EAVPackage }}"
    {{ range .ImportPaths }} "{{.}}"
    {{ end }} )

{{.TypeDefinition}}

const (
    {{ range $k, $row := .Attributes }}{{$.Name | prepareVar}}{{index $row "attribute_code" | prepareVar}} {{ if eq $k 0 }} eav.AttributeIndex = iota {{ end }}
    {{end}}
    {{$.Name | prepareVar}}999Max
)

type si{{$.Name | prepareVar}} struct {}

func (si{{$.Name | prepareVar}}) ByID(id int64) (eav.AttributeIndex, error){
	switch id {
	{{ range $k, $row := .Attributes }} case {{index $row "attribute_id"}}:
		return {{$.Name | prepareVar}}{{index $row "attribute_code" | prepareVar}}, nil
	{{end}}
	default:
		return eav.AttributeIndex(0), eav.ErrAttributeNotFound
	}
}

func (si{{$.Name | prepareVar}}) ByCode(code string) (eav.AttributeIndex, error){
	switch code {
	{{ range $k, $row := .Attributes }} case {{index $row "attribute_code"}}:
		return {{$.Name | prepareVar}}{{index $row "attribute_code" | prepareVar}}, nil
	{{end}}
	default:
		return eav.AttributeIndex(0), eav.ErrAttributeNotFound
	}
}

var _ eav.AttributeGetter = (*si{{$.Name | prepareVar}})(nil)

func init(){
    {{.EAVPackage}}.SetAttributeCollection({{.EAVPackage}}.AttributeSlice{
        {{ range $row := .Attributes }} {{$.Name | prepareVar}}{{index $row "attribute_code" | prepareVar}}: &{{$.Name | prepareVar | toLowerFirst}} {
            {{ range $k,$v := $row }} {{ $k | prepareVar | toLowerFirst }}: {{ $v }},
            {{ end }}
        },
        {{ end }}
    })
}
`
