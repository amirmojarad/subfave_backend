// Code generated by entc, DO NOT EDIT.

package word

const (
	// Label holds the string label denoting the word type in the database.
	Label = "word"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldMeaning holds the string denoting the meaning field in the database.
	FieldMeaning = "meaning"
	// FieldSentence holds the string denoting the sentence field in the database.
	FieldSentence = "sentence"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// Table holds the table name of the word in the database.
	Table = "words"
)

// Columns holds all SQL columns for word fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldMeaning,
	FieldSentence,
	FieldDuration,
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
)
