// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MoviesColumns holds the columns for the "movies" table.
	MoviesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "year", Type: field.TypeString},
		{Name: "image_url", Type: field.TypeString},
		{Name: "runtime_str", Type: field.TypeString},
		{Name: "genres", Type: field.TypeString},
		{Name: "im_db_rating", Type: field.TypeString},
		{Name: "plot", Type: field.TypeString},
		{Name: "stars", Type: field.TypeString},
		{Name: "metacritic_rating", Type: field.TypeString},
		{Name: "word_node_movie_wordnode", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// MoviesTable holds the schema information for the "movies" table.
	MoviesTable = &schema.Table{
		Name:       "movies",
		Columns:    MoviesColumns,
		PrimaryKey: []*schema.Column{MoviesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "movies_word_nodes_movie_wordnode",
				Columns:    []*schema.Column{MoviesColumns[10]},
				RefColumns: []*schema.Column{WordNodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SearchKeywordsColumns holds the columns for the "search_keywords" table.
	SearchKeywordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "rate", Type: field.TypeUint16, Default: 0},
		{Name: "user_searched_keywords", Type: field.TypeInt, Nullable: true},
	}
	// SearchKeywordsTable holds the schema information for the "search_keywords" table.
	SearchKeywordsTable = &schema.Table{
		Name:       "search_keywords",
		Columns:    SearchKeywordsColumns,
		PrimaryKey: []*schema.Column{SearchKeywordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "search_keywords_users_searched_keywords",
				Columns:    []*schema.Column{SearchKeywordsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "full_name", Type: field.TypeString, Nullable: true, Default: ""},
		{Name: "created_date", Type: field.TypeTime},
		{Name: "updated_date", Type: field.TypeTime},
		{Name: "is_admin", Type: field.TypeBool, Default: false},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// WordsColumns holds the columns for the "words" table.
	WordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "meaning", Type: field.TypeString, Nullable: true},
		{Name: "sentence", Type: field.TypeString, Nullable: true},
		{Name: "duration", Type: field.TypeString, Nullable: true},
		{Name: "start", Type: field.TypeTime, Nullable: true},
		{Name: "end", Type: field.TypeTime, Nullable: true},
		{Name: "user_favorite_words", Type: field.TypeInt, Nullable: true},
		{Name: "word_movie", Type: field.TypeInt, Nullable: true},
		{Name: "word_node_words", Type: field.TypeInt, Nullable: true},
	}
	// WordsTable holds the schema information for the "words" table.
	WordsTable = &schema.Table{
		Name:       "words",
		Columns:    WordsColumns,
		PrimaryKey: []*schema.Column{WordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "words_users_favorite_words",
				Columns:    []*schema.Column{WordsColumns[7]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "words_movies_movie",
				Columns:    []*schema.Column{WordsColumns[8]},
				RefColumns: []*schema.Column{MoviesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "words_word_nodes_words",
				Columns:    []*schema.Column{WordsColumns[9]},
				RefColumns: []*schema.Column{WordNodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// WordNodesColumns holds the columns for the "word_nodes" table.
	WordNodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "is_preposition", Type: field.TypeBool},
		{Name: "occurence", Type: field.TypeInt, Nullable: true},
	}
	// WordNodesTable holds the schema information for the "word_nodes" table.
	WordNodesTable = &schema.Table{
		Name:       "word_nodes",
		Columns:    WordNodesColumns,
		PrimaryKey: []*schema.Column{WordNodesColumns[0]},
	}
	// UserFavoriteMoviesColumns holds the columns for the "user_favorite_movies" table.
	UserFavoriteMoviesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "movie_id", Type: field.TypeInt},
	}
	// UserFavoriteMoviesTable holds the schema information for the "user_favorite_movies" table.
	UserFavoriteMoviesTable = &schema.Table{
		Name:       "user_favorite_movies",
		Columns:    UserFavoriteMoviesColumns,
		PrimaryKey: []*schema.Column{UserFavoriteMoviesColumns[0], UserFavoriteMoviesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_favorite_movies_user_id",
				Columns:    []*schema.Column{UserFavoriteMoviesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_favorite_movies_movie_id",
				Columns:    []*schema.Column{UserFavoriteMoviesColumns[1]},
				RefColumns: []*schema.Column{MoviesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MoviesTable,
		SearchKeywordsTable,
		UsersTable,
		WordsTable,
		WordNodesTable,
		UserFavoriteMoviesTable,
	}
)

func init() {
	MoviesTable.ForeignKeys[0].RefTable = WordNodesTable
	SearchKeywordsTable.ForeignKeys[0].RefTable = UsersTable
	WordsTable.ForeignKeys[0].RefTable = UsersTable
	WordsTable.ForeignKeys[1].RefTable = MoviesTable
	WordsTable.ForeignKeys[2].RefTable = WordNodesTable
	UserFavoriteMoviesTable.ForeignKeys[0].RefTable = UsersTable
	UserFavoriteMoviesTable.ForeignKeys[1].RefTable = MoviesTable
}
