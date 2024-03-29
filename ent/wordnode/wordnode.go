// Code generated by entc, DO NOT EDIT.

package wordnode

const (
	// Label holds the string label denoting the wordnode type in the database.
	Label = "word_node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// EdgeWord holds the string denoting the word edge name in mutations.
	EdgeWord = "word"
	// EdgeSentence holds the string denoting the sentence edge name in mutations.
	EdgeSentence = "sentence"
	// Table holds the table name of the wordnode in the database.
	Table = "word_nodes"
	// WordTable is the table that holds the word relation/edge.
	WordTable = "words"
	// WordInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	WordInverseTable = "words"
	// WordColumn is the table column denoting the word relation/edge.
	WordColumn = "word_node_word"
	// SentenceTable is the table that holds the sentence relation/edge.
	SentenceTable = "sentences"
	// SentenceInverseTable is the table name for the Sentence entity.
	// It exists in this package in order to avoid circular dependency with the "sentence" package.
	SentenceInverseTable = "sentences"
	// SentenceColumn is the table column denoting the sentence relation/edge.
	SentenceColumn = "word_node_sentence"
)

// Columns holds all SQL columns for wordnode fields.
var Columns = []string{
	FieldID,
	FieldDuration,
	FieldStart,
	FieldEnd,
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
	// DurationValidator is a validator for the "duration" field. It is called by the builders before save.
	DurationValidator func(string) error
)
