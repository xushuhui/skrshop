// Code generated by entc, DO NOT EDIT.

package authms

const (
	// Label holds the string label denoting the authms type in the database.
	Label = "auth_ms"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"

	// Table holds the table name of the authms in the database.
	Table = "auth_ms"
)

// Columns holds all SQL columns for authms fields.
var Columns = []string{
	FieldID,
	FieldStatus,
	FieldCreateAt,
	FieldUpdateAt,
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
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)
