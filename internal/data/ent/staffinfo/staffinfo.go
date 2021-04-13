// Code generated by entc, DO NOT EDIT.

package staffinfo

const (
	// Label holds the string label denoting the staffinfo type in the database.
	Label = "staff_info"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUID holds the string denoting the uid field in the database.
	FieldUID = "uid"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"

	// Table holds the table name of the staffinfo in the database.
	Table = "staff_infos"
)

// Columns holds all SQL columns for staffinfo fields.
var Columns = []string{
	FieldID,
	FieldUID,
	FieldEmail,
	FieldPhone,
	FieldPassword,
	FieldName,
	FieldAvatar,
	FieldGender,
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
	// DefaultUID holds the default value on creation for the "uid" field.
	DefaultUID int64
	// UIDValidator is a validator for the "uid" field. It is called by the builders before save.
	UIDValidator func(int64) error
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
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultAvatar holds the default value on creation for the "avatar" field.
	DefaultAvatar string
	// AvatarValidator is a validator for the "avatar" field. It is called by the builders before save.
	AvatarValidator func(string) error
	// DefaultGender holds the default value on creation for the "gender" field.
	DefaultGender int8
)
