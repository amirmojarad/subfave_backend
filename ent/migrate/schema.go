// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CollectionsColumns holds the columns for the "collections" table.
	CollectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "created_date", Type: field.TypeTime},
		{Name: "updated_date", Type: field.TypeTime},
	}
	// CollectionsTable holds the schema information for the "collections" table.
	CollectionsTable = &schema.Table{
		Name:       "collections",
		Columns:    CollectionsColumns,
		PrimaryKey: []*schema.Column{CollectionsColumns[0]},
	}
	// FileEntitiesColumns holds the columns for the "file_entities" table.
	FileEntitiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "path", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "size", Type: field.TypeInt16},
		{Name: "deleted", Type: field.TypeBool},
		{Name: "created_date", Type: field.TypeTime},
		{Name: "user_files", Type: field.TypeInt, Nullable: true},
	}
	// FileEntitiesTable holds the schema information for the "file_entities" table.
	FileEntitiesTable = &schema.Table{
		Name:       "file_entities",
		Columns:    FileEntitiesColumns,
		PrimaryKey: []*schema.Column{FileEntitiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "file_entities_users_files",
				Columns:    []*schema.Column{FileEntitiesColumns[6]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
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
	}
	// MoviesTable holds the schema information for the "movies" table.
	MoviesTable = &schema.Table{
		Name:       "movies",
		Columns:    MoviesColumns,
		PrimaryKey: []*schema.Column{MoviesColumns[0]},
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
				OnDelete:   schema.Cascade,
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
		{Name: "image_url", Type: field.TypeString, Nullable: true},
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
		{Name: "meaning", Type: field.TypeString},
		{Name: "is_preposition", Type: field.TypeBool},
		{Name: "sentence", Type: field.TypeString},
		{Name: "duration", Type: field.TypeString},
		{Name: "start", Type: field.TypeString},
		{Name: "end", Type: field.TypeString},
		{Name: "file_entity_words", Type: field.TypeInt, Nullable: true},
		{Name: "user_favorite_words", Type: field.TypeInt, Nullable: true},
		{Name: "user_words", Type: field.TypeInt, Nullable: true},
	}
	// WordsTable holds the schema information for the "words" table.
	WordsTable = &schema.Table{
		Name:       "words",
		Columns:    WordsColumns,
		PrimaryKey: []*schema.Column{WordsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "words_file_entities_words",
				Columns:    []*schema.Column{WordsColumns[8]},
				RefColumns: []*schema.Column{FileEntitiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "words_users_favorite_words",
				Columns:    []*schema.Column{WordsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "words_users_words",
				Columns:    []*schema.Column{WordsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// CollectionCollectionWordsColumns holds the columns for the "collection_collection_words" table.
	CollectionCollectionWordsColumns = []*schema.Column{
		{Name: "collection_id", Type: field.TypeInt},
		{Name: "word_id", Type: field.TypeInt},
	}
	// CollectionCollectionWordsTable holds the schema information for the "collection_collection_words" table.
	CollectionCollectionWordsTable = &schema.Table{
		Name:       "collection_collection_words",
		Columns:    CollectionCollectionWordsColumns,
		PrimaryKey: []*schema.Column{CollectionCollectionWordsColumns[0], CollectionCollectionWordsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "collection_collection_words_collection_id",
				Columns:    []*schema.Column{CollectionCollectionWordsColumns[0]},
				RefColumns: []*schema.Column{CollectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "collection_collection_words_word_id",
				Columns:    []*schema.Column{CollectionCollectionWordsColumns[1]},
				RefColumns: []*schema.Column{WordsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
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
	// UserCollectionsColumns holds the columns for the "user_collections" table.
	UserCollectionsColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "collection_id", Type: field.TypeInt},
	}
	// UserCollectionsTable holds the schema information for the "user_collections" table.
	UserCollectionsTable = &schema.Table{
		Name:       "user_collections",
		Columns:    UserCollectionsColumns,
		PrimaryKey: []*schema.Column{UserCollectionsColumns[0], UserCollectionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_collections_user_id",
				Columns:    []*schema.Column{UserCollectionsColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_collections_collection_id",
				Columns:    []*schema.Column{UserCollectionsColumns[1]},
				RefColumns: []*schema.Column{CollectionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CollectionsTable,
		FileEntitiesTable,
		MoviesTable,
		SearchKeywordsTable,
		UsersTable,
		WordsTable,
		CollectionCollectionWordsTable,
		UserFavoriteMoviesTable,
		UserCollectionsTable,
	}
)

func init() {
	FileEntitiesTable.ForeignKeys[0].RefTable = UsersTable
	SearchKeywordsTable.ForeignKeys[0].RefTable = UsersTable
	WordsTable.ForeignKeys[0].RefTable = FileEntitiesTable
	WordsTable.ForeignKeys[1].RefTable = UsersTable
	WordsTable.ForeignKeys[2].RefTable = UsersTable
	CollectionCollectionWordsTable.ForeignKeys[0].RefTable = CollectionsTable
	CollectionCollectionWordsTable.ForeignKeys[1].RefTable = WordsTable
	UserFavoriteMoviesTable.ForeignKeys[0].RefTable = UsersTable
	UserFavoriteMoviesTable.ForeignKeys[1].RefTable = MoviesTable
	UserCollectionsTable.ForeignKeys[0].RefTable = UsersTable
	UserCollectionsTable.ForeignKeys[1].RefTable = CollectionsTable
}
