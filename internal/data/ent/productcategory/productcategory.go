// Code generated by entc, DO NOT EDIT.

package productcategory

const (
	// Label holds the string label denoting the productcategory type in the database.
	Label = "product_category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPid holds the string denoting the pid field in the database.
	FieldPid = "pid"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldPicURL holds the string denoting the pic_url field in the database.
	FieldPicURL = "pic_url"
	// FieldPath holds the string denoting the path field in the database.
	FieldPath = "path"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdateAt holds the string denoting the update_at field in the database.
	FieldUpdateAt = "update_at"

	// Table holds the table name of the productcategory in the database.
	Table = "product_categories"
)

// Columns holds all SQL columns for productcategory fields.
var Columns = []string{
	FieldID,
	FieldPid,
	FieldName,
	FieldDesc,
	FieldPicURL,
	FieldPath,
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
	// DefaultPid holds the default value on creation for the "pid" field.
	DefaultPid int64
	// PidValidator is a validator for the "pid" field. It is called by the builders before save.
	PidValidator func(int64) error
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultDesc holds the default value on creation for the "desc" field.
	DefaultDesc string
	// DescValidator is a validator for the "desc" field. It is called by the builders before save.
	DescValidator func(string) error
	// DefaultPicURL holds the default value on creation for the "pic_url" field.
	DefaultPicURL string
	// PicURLValidator is a validator for the "pic_url" field. It is called by the builders before save.
	PicURLValidator func(string) error
	// DefaultPath holds the default value on creation for the "path" field.
	DefaultPath string
	// PathValidator is a validator for the "path" field. It is called by the builders before save.
	PathValidator func(string) error
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus int8
)
