// Code generated by entc, DO NOT EDIT.

package movie

const (
	// Label holds the string label denoting the movie type in the database.
	Label = "movie"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldYear holds the string denoting the year field in the database.
	FieldYear = "year"
	// FieldImageURL holds the string denoting the image_url field in the database.
	FieldImageURL = "image_url"
	// FieldRuntimeStr holds the string denoting the runtimestr field in the database.
	FieldRuntimeStr = "runtime_str"
	// FieldGenres holds the string denoting the genres field in the database.
	FieldGenres = "genres"
	// FieldImDbRating holds the string denoting the imdbrating field in the database.
	FieldImDbRating = "im_db_rating"
	// FieldPlot holds the string denoting the plot field in the database.
	FieldPlot = "plot"
	// FieldStars holds the string denoting the stars field in the database.
	FieldStars = "stars"
	// FieldMetacriticRating holds the string denoting the metacriticrating field in the database.
	FieldMetacriticRating = "metacritic_rating"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// Table holds the table name of the movie in the database.
	Table = "movies"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_favorite_movies"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
)

// Columns holds all SQL columns for movie fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldYear,
	FieldImageURL,
	FieldRuntimeStr,
	FieldGenres,
	FieldImDbRating,
	FieldPlot,
	FieldStars,
	FieldMetacriticRating,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "movie_id"}
)

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
