// Code generated by entc, DO NOT EDIT.

package category

import (
	"time"
)

const (
	// Label holds the string label denoting the category type in the database.
	Label = "category"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// FieldUpdatedDate holds the string denoting the updated_date field in the database.
	FieldUpdatedDate = "updated_date"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeCategoryWords holds the string denoting the category_words edge name in mutations.
	EdgeCategoryWords = "category_words"
	// Table holds the table name of the category in the database.
	Table = "categories"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "user_categories"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// CategoryWordsTable is the table that holds the category_words relation/edge. The primary key declared below.
	CategoryWordsTable = "category_category_words"
	// CategoryWordsInverseTable is the table name for the Word entity.
	// It exists in this package in order to avoid circular dependency with the "word" package.
	CategoryWordsInverseTable = "words"
)

// Columns holds all SQL columns for category fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldCreatedDate,
	FieldUpdatedDate,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "category_id"}
	// CategoryWordsPrimaryKey and CategoryWordsColumn2 are the table columns denoting the
	// primary key for the category_words relation (M2M).
	CategoryWordsPrimaryKey = []string{"category_id", "word_id"}
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
	// DefaultCreatedDate holds the default value on creation for the "created_date" field.
	DefaultCreatedDate time.Time
	// DefaultUpdatedDate holds the default value on creation for the "updated_date" field.
	DefaultUpdatedDate time.Time
)
