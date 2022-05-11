// Code generated by entc, DO NOT EDIT.

package wordnode

const (
	// Label holds the string label denoting the wordnode type in the database.
	Label = "word_node"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldIsPreposition holds the string denoting the is_preposition field in the database.
	FieldIsPreposition = "is_preposition"
	// FieldOccurence holds the string denoting the occurence field in the database.
	FieldOccurence = "occurence"
	// EdgeWords holds the string denoting the words edge name in mutations.
	EdgeWords = "words"
	// EdgeMovieWordnode holds the string denoting the movie_wordnode edge name in mutations.
	EdgeMovieWordnode = "movie_wordnode"
	// Table holds the table name of the wordnode in the database.
	Table = "word_nodes"
	// WordsTable is the table that holds the words relation/edge.
	WordsTable = "words"
	// WordsInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	WordsInverseTable = "words"
	// WordsColumn is the table column denoting the words relation/edge.
	WordsColumn = "word_node_words"
	// MovieWordnodeTable is the table that holds the movie_wordnode relation/edge.
	MovieWordnodeTable = "movies"
	// MovieWordnodeInverseTable is the table name for the Movie entity.
	// It exists in this package in order to avoid circular dependency with the "movie" package.
	MovieWordnodeInverseTable = "movies"
	// MovieWordnodeColumn is the table column denoting the movie_wordnode relation/edge.
	MovieWordnodeColumn = "word_node_movie_wordnode"
)

// Columns holds all SQL columns for wordnode fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldIsPreposition,
	FieldOccurence,
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
