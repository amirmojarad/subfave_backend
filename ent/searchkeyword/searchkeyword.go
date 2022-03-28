// Code generated by entc, DO NOT EDIT.

package searchkeyword

const (
	// Label holds the string label denoting the searchkeyword type in the database.
	Label = "search_keyword"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldRate holds the string denoting the rate field in the database.
	FieldRate = "rate"
	// Table holds the table name of the searchkeyword in the database.
	Table = "search_keywords"
)

// Columns holds all SQL columns for searchkeyword fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldRate,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "search_keywords"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_searched_keywords",
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DefaultRate holds the default value on creation for the "rate" field.
	DefaultRate uint16
)
