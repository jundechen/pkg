// Code generated by codegen. DO NOT EDIT.
// Generated by sql/dmlgen. DO NOT EDIT.
package dmltestgenerated2

import (
	"fmt"
	"github.com/corestoreio/errors"
	"github.com/corestoreio/pkg/storage/null"
	"io"
	"time"
)

// CoreConfiguration represents a single row for DB table core_configuration.
// Auto generated.
type CoreConfiguration struct {
	ConfigID  uint32      // config_id int(10) unsigned NOT NULL PRI  auto_increment "Id"
	Scope     string      // scope varchar(8) NOT NULL MUL DEFAULT ''default''  "Scope"
	ScopeID   int32       // scope_id int(11) NOT NULL  DEFAULT '0'  "Scope Id"
	Expires   null.Time   // expires datetime NULL  DEFAULT 'NULL'  "Value expiration time"
	Path      string      // path varchar(255) NOT NULL    "Path"
	Value     null.String // value text NULL  DEFAULT 'NULL'  "Value"
	VersionTs time.Time   // version_ts timestamp(6) NOT NULL    "Timestamp Start Versioning"
	VersionTe time.Time   // version_te timestamp(6) NOT NULL PRI   "Timestamp End Versioning"
}

// Empty empties all the fields of the current object. Also known as Reset.
func (e *CoreConfiguration) Empty() *CoreConfiguration { *e = CoreConfiguration{}; return e }

// Copy copies the struct and returns a new pointer
func (e *CoreConfiguration) Copy() *CoreConfiguration {
	e2 := new(CoreConfiguration)
	*e2 = *e // for now a shallow copy
	return e2
}

// WriteTo implements io.WriterTo and writes the field names and their values to
// w. This is especially useful for debugging or or generating a hash of the
// struct.
func (e *CoreConfiguration) WriteTo(w io.Writer) (n int64, err error) {
	// for now this printing is good enough. If you need better swap out with your code.
	n2, err := fmt.Fprint(w,
		"config_id:", e.ConfigID, "\n",
		"scope:", e.Scope, "\n",
		"scope_id:", e.ScopeID, "\n",
		"expires:", e.Expires, "\n",
		"path:", e.Path, "\n",
		"value:", e.Value, "\n",
		"version_ts:", e.VersionTs, "\n",
		"version_te:", e.VersionTe, "\n",
	)
	return int64(n2), err
}

// CoreConfigurationCollection represents a collection type for DB table
// core_configuration
// Not thread safe. Auto generated.
type CoreConfigurationCollection struct {
	Data []*CoreConfiguration `json:"data,omitempty"`
}

// NewCoreConfigurationCollection  creates a new initialized collection. Auto
// generated.
func NewCoreConfigurationCollection() *CoreConfigurationCollection {
	return &CoreConfigurationCollection{
		Data: make([]*CoreConfiguration, 0, 5),
	}
}

// ConfigIDs returns a slice with the data or appends it to a slice.
// Auto generated.
func (cc *CoreConfigurationCollection) ConfigIDs(ret ...uint32) []uint32 {
	if ret == nil {
		ret = make([]uint32, 0, len(cc.Data))
	}
	for _, e := range cc.Data {
		ret = append(ret, e.ConfigID)
	}
	return ret
}

// WriteTo implements io.WriterTo and writes the field names and their values to
// w. This is especially useful for debugging or or generating a hash of the
// struct.
func (cc *CoreConfigurationCollection) WriteTo(w io.Writer) (n int64, err error) {
	for i, d := range cc.Data {
		n2, err := d.WriteTo(w)
		if err != nil {
			return 0, errors.Wrapf(err, "[dmltestgenerated2] WriteTo failed at index %d", i)
		}
		n += n2
	}
	return n, nil
}

// Filter filters the current slice by predicate f without memory allocation.
// Auto generated via dmlgen.
func (cc *CoreConfigurationCollection) Filter(f func(*CoreConfiguration) bool) *CoreConfigurationCollection {
	b, i := cc.Data[:0], 0
	for _, e := range cc.Data {
		if f(e) {
			b = append(b, e)
			cc.Data[i] = nil // this avoids the memory leak
		}
		i++
	}
	cc.Data = b
	return cc
}

// Each will run function f on all items in []* CoreConfiguration . Auto
// generated via dmlgen.
func (cc *CoreConfigurationCollection) Each(f func(*CoreConfiguration)) *CoreConfigurationCollection {
	for i := range cc.Data {
		f(cc.Data[i])
	}
	return cc
}

// Cut will remove items i through j-1. Auto generated via dmlgen.
func (cc *CoreConfigurationCollection) Cut(i, j int) *CoreConfigurationCollection {
	z := cc.Data // copy slice header
	copy(z[i:], z[j:])
	for k, n := len(z)-j+i, len(z); k < n; k++ {
		z[k] = nil // this avoids the memory leak
	}
	z = z[:len(z)-j+i]
	cc.Data = z
	return cc
}

// Swap will satisfy the sort.Interface. Auto generated via dmlgen.
func (cc *CoreConfigurationCollection) Swap(i, j int) {
	cc.Data[i], cc.Data[j] = cc.Data[j], cc.Data[i]
}

// Len will satisfy the sort.Interface. Auto generated via dmlgen.
func (cc *CoreConfigurationCollection) Len() int { return len(cc.Data) }

// Delete will remove an item from the slice. Auto generated via dmlgen.
func (cc *CoreConfigurationCollection) Delete(i int) *CoreConfigurationCollection {
	z := cc.Data // copy the slice header
	end := len(z) - 1
	cc.Swap(i, end)
	copy(z[i:], z[i+1:])
	z[end] = nil // this should avoid the memory leak
	z = z[:end]
	cc.Data = z
	return cc
}

// Insert will place a new item at position i. Auto generated via dmlgen.
func (cc *CoreConfigurationCollection) Insert(n *CoreConfiguration, i int) *CoreConfigurationCollection {
	z := cc.Data // copy the slice header
	z = append(z, &CoreConfiguration{})
	copy(z[i+1:], z[i:])
	z[i] = n
	cc.Data = z
	return cc
}

// Append will add a new item at the end of * CoreConfigurationCollection . Auto
// generated via dmlgen.
func (cc *CoreConfigurationCollection) Append(n ...*CoreConfiguration) *CoreConfigurationCollection {
	cc.Data = append(cc.Data, n...)
	return cc
}

// SalesOrderStatusState represents a single row for DB table
// sales_order_status_state. Auto generated.
type SalesOrderStatusState struct {
	Status         string // status varchar(32) NOT NULL PRI   "Status"
	State          string // state varchar(32) NOT NULL PRI   "Label"
	IsDefault      bool   // is_default smallint(5) unsigned NOT NULL  DEFAULT '0'  "Is Default"
	VisibleOnFront uint16 // visible_on_front smallint(5) unsigned NOT NULL  DEFAULT '0'  "Visible on front"
}

// Empty empties all the fields of the current object. Also known as Reset.
func (e *SalesOrderStatusState) Empty() *SalesOrderStatusState {
	*e = SalesOrderStatusState{}
	return e
}

// Copy copies the struct and returns a new pointer
func (e *SalesOrderStatusState) Copy() *SalesOrderStatusState {
	e2 := new(SalesOrderStatusState)
	*e2 = *e // for now a shallow copy
	return e2
}

// WriteTo implements io.WriterTo and writes the field names and their values to
// w. This is especially useful for debugging or or generating a hash of the
// struct.
func (e *SalesOrderStatusState) WriteTo(w io.Writer) (n int64, err error) {
	// for now this printing is good enough. If you need better swap out with your code.
	n2, err := fmt.Fprint(w,
		"status:", e.Status, "\n",
		"state:", e.State, "\n",
		"is_default:", e.IsDefault, "\n",
		"visible_on_front:", e.VisibleOnFront, "\n",
	)
	return int64(n2), err
}

// SalesOrderStatusStateCollection represents a collection type for DB table
// sales_order_status_state
// Not thread safe. Auto generated.
type SalesOrderStatusStateCollection struct {
	Data []*SalesOrderStatusState `json:"data,omitempty"`
}

// NewSalesOrderStatusStateCollection  creates a new initialized collection. Auto
// generated.
func NewSalesOrderStatusStateCollection() *SalesOrderStatusStateCollection {
	return &SalesOrderStatusStateCollection{
		Data: make([]*SalesOrderStatusState, 0, 5),
	}
}

// Statuss returns a slice with the data or appends it to a slice.
// Auto generated.
func (cc *SalesOrderStatusStateCollection) Statuss(ret ...string) []string {
	if ret == nil {
		ret = make([]string, 0, len(cc.Data))
	}
	for _, e := range cc.Data {
		ret = append(ret, e.Status)
	}
	return ret
}

// States returns a slice with the data or appends it to a slice.
// Auto generated.
func (cc *SalesOrderStatusStateCollection) States(ret ...string) []string {
	if ret == nil {
		ret = make([]string, 0, len(cc.Data))
	}
	for _, e := range cc.Data {
		ret = append(ret, e.State)
	}
	return ret
}

// WriteTo implements io.WriterTo and writes the field names and their values to
// w. This is especially useful for debugging or or generating a hash of the
// struct.
func (cc *SalesOrderStatusStateCollection) WriteTo(w io.Writer) (n int64, err error) {
	for i, d := range cc.Data {
		n2, err := d.WriteTo(w)
		if err != nil {
			return 0, errors.Wrapf(err, "[dmltestgenerated2] WriteTo failed at index %d", i)
		}
		n += n2
	}
	return n, nil
}

// Filter filters the current slice by predicate f without memory allocation.
// Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Filter(f func(*SalesOrderStatusState) bool) *SalesOrderStatusStateCollection {
	b, i := cc.Data[:0], 0
	for _, e := range cc.Data {
		if f(e) {
			b = append(b, e)
			cc.Data[i] = nil // this avoids the memory leak
		}
		i++
	}
	cc.Data = b
	return cc
}

// Each will run function f on all items in []* SalesOrderStatusState . Auto
// generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Each(f func(*SalesOrderStatusState)) *SalesOrderStatusStateCollection {
	for i := range cc.Data {
		f(cc.Data[i])
	}
	return cc
}

// Cut will remove items i through j-1. Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Cut(i, j int) *SalesOrderStatusStateCollection {
	z := cc.Data // copy slice header
	copy(z[i:], z[j:])
	for k, n := len(z)-j+i, len(z); k < n; k++ {
		z[k] = nil // this avoids the memory leak
	}
	z = z[:len(z)-j+i]
	cc.Data = z
	return cc
}

// Swap will satisfy the sort.Interface. Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Swap(i, j int) {
	cc.Data[i], cc.Data[j] = cc.Data[j], cc.Data[i]
}

// Len will satisfy the sort.Interface. Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Len() int { return len(cc.Data) }

// Delete will remove an item from the slice. Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Delete(i int) *SalesOrderStatusStateCollection {
	z := cc.Data // copy the slice header
	end := len(z) - 1
	cc.Swap(i, end)
	copy(z[i:], z[i+1:])
	z[end] = nil // this should avoid the memory leak
	z = z[:end]
	cc.Data = z
	return cc
}

// Insert will place a new item at position i. Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Insert(n *SalesOrderStatusState, i int) *SalesOrderStatusStateCollection {
	z := cc.Data // copy the slice header
	z = append(z, &SalesOrderStatusState{})
	copy(z[i+1:], z[i:])
	z[i] = n
	cc.Data = z
	return cc
}

// Append will add a new item at the end of * SalesOrderStatusStateCollection .
// Auto generated via dmlgen.
func (cc *SalesOrderStatusStateCollection) Append(n ...*SalesOrderStatusState) *SalesOrderStatusStateCollection {
	cc.Data = append(cc.Data, n...)
	return cc
}
