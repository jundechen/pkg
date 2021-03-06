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

package customer

import (
	"github.com/corestoreio/pkg/eav"
	"github.com/corestoreio/pkg/storage/csdb"
)

type (
	// see table customer_eav_attribute.data_model
	DataModeller interface {
		TBD()
	}
	// Entity is the customer model
	Entity struct {
		// TBD
	}
)

var (
	_ eav.EntityTypeModeller                  = (*Entity)(nil)
	_ eav.EntityTypeTabler                    = (*Entity)(nil)
	_ eav.EntityTypeAdditionalAttributeTabler = (*Entity)(nil)
	_ eav.EntityTypeIncrementModeller         = (*Entity)(nil)
	// TableCollection handles all tables and its columns. init() in generated Go file will set the value.
	TableCollection csdb.TableManager
)

func (c *Entity) TBD() {

}

func (c *Entity) TableNameBase() string {
	return TableCollection.Name(TableIndexEntity)
}

func (c *Entity) TableNameValue(i eav.ValueIndex) string {
	s, err := GetCustomerValueStructure(i)
	if err != nil {
		return ""
	}
	return s.Name
}

// EntityTypeAdditionalAttributeTabler
func (c *Entity) TableAdditionalAttribute() (*csdb.Table, error) {
	return TableCollection.Structure(TableIndexEAVAttribute)
}

// EntityTypeAdditionalAttributeTabler
func (c *Entity) TableEavWebsite() (*csdb.Table, error) {
	return TableCollection.Structure(TableIndexEAVAttributeWebsite)
}

func Customer() *Entity {
	return &Entity{}
}
