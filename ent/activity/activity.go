// Code generated by ent, DO NOT EDIT.

package activity

const (
	// Label holds the string label denoting the activity type in the database.
	Label = "activity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "activity"
	// Table holds the table name of the activity in the database.
	Table = "activities"
)

// Columns holds all SQL columns for activity fields.
var Columns = []string{
	FieldID,
	FieldUsername,
	FieldDuration,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "activities"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"day_activity",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDuration holds the default value on creation for the "duration" field.
	DefaultDuration int64
)
