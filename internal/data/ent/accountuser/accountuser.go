// Code generated by entc, DO NOT EDIT.

package accountuser

const (
	// Label holds the string label denoting the accountuser type in the database.
	Label = "account_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"

	// Table holds the table name of the accountuser in the database.
	Table = "account_users"
)

// Columns holds all SQL columns for accountuser fields.
var Columns = []string{
	FieldID,
	FieldEmail,
	FieldPhone,
	FieldPassword,
	FieldCreateAt,
	FieldStatus,
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
	// DefaultEmail holds the default value on creation for the "email" field.
	DefaultEmail string
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultPhone holds the default value on creation for the "phone" field.
	DefaultPhone string
	// PhoneValidator is a validator for the "phone" field. It is called by the builders before save.
	PhoneValidator func(string) error
	// DefaultPassword holds the default value on creation for the "password" field.
	DefaultPassword string
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)
