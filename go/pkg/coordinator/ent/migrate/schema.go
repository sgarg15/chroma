// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BasesColumns holds the columns for the "bases" table.
	BasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "parent_id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Nullable: true, Size: 255},
		{Name: "created_at", Type: field.TypeUint},
		{Name: "updated_at", Type: field.TypeUint},
		{Name: "deleted_at", Type: field.TypeUint, Default: 0},
		{Name: "version", Type: field.TypeInt, Default: 0},
	}
	// BasesTable holds the schema information for the "bases" table.
	BasesTable = &schema.Table{
		Name:       "bases",
		Columns:    BasesColumns,
		PrimaryKey: []*schema.Column{BasesColumns[0]},
	}
	// TestBasesColumns holds the columns for the "test_bases" table.
	TestBasesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// TestBasesTable holds the schema information for the "test_bases" table.
	TestBasesTable = &schema.Table{
		Name:       "test_bases",
		Columns:    TestBasesColumns,
		PrimaryKey: []*schema.Column{TestBasesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BasesTable,
		TestBasesTable,
	}
)

func init() {
}
