// Code generated by ent, DO NOT EDIT.

package base

import (
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the base type in the database.
	Label = "base"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldParentID holds the string denoting the parent_id field in the database.
	FieldParentID = "parent_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldVersion holds the string denoting the version field in the database.
	FieldVersion = "version"
	// Table holds the table name of the base in the database.
	Table = "bases"
)

// Columns holds all SQL columns for base fields.
var Columns = []string{
	FieldID,
	FieldParentID,
	FieldName,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldDeletedAt,
	FieldVersion,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() uint
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() uint
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt uint
	// DefaultVersion holds the default value on creation for the "version" field.
	DefaultVersion int
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// OrderOption defines the ordering options for the Base queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByParentID orders the results by the parent_id field.
func ByParentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldParentID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByVersion orders the results by the version field.
func ByVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVersion, opts...).ToFunc()
}
