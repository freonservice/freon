// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package dao

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type categoryTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *categoryTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("categories").
func (v *categoryTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *categoryTableType) Columns() []string {
	return []string{
		"id",
		"name",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *categoryTableType) NewStruct() reform.Struct {
	return new(Category)
}

// NewRecord makes a new record for that table.
func (v *categoryTableType) NewRecord() reform.Record {
	return new(Category)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *categoryTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// CategoryTable represents categories view or table in SQL database.
var CategoryTable = &categoryTableType{
	s: parse.StructInfo{
		Type:    "Category",
		SQLName: "categories",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int64", Column: "id"},
			{Name: "Name", Type: "string", Column: "name"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "*time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(Category).Values(),
}

// String returns a string representation of this struct or record.
func (s Category) String() string {
	res := make([]string, 4)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Name: " + reform.Inspect(s.Name, true)
	res[2] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[3] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *Category) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Name,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *Category) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Name,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *Category) View() reform.View {
	return CategoryTable
}

// Table returns Table object for that record.
func (s *Category) Table() reform.Table {
	return CategoryTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *Category) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *Category) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *Category) HasPK() bool {
	return s.ID != CategoryTable.z[CategoryTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *Category) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = CategoryTable
	_ reform.Struct = (*Category)(nil)
	_ reform.Table  = CategoryTable
	_ reform.Record = (*Category)(nil)
	_ fmt.Stringer  = (*Category)(nil)
)

func init() {
	parse.AssertUpToDate(&CategoryTable.s, new(Category))
}
