// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"interface_project/ent/searchkeyword"
	"interface_project/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// SearchKeyword is the model entity for the SearchKeyword schema.
type SearchKeyword struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Rate holds the value of the "rate" field.
	Rate uint16 `json:"rate,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SearchKeywordQuery when eager-loading is set.
	Edges                  SearchKeywordEdges `json:"edges"`
	user_searched_keywords *int
}

// SearchKeywordEdges holds the relations/edges for other nodes in the graph.
type SearchKeywordEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SearchKeywordEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SearchKeyword) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case searchkeyword.FieldID, searchkeyword.FieldRate:
			values[i] = new(sql.NullInt64)
		case searchkeyword.FieldTitle:
			values[i] = new(sql.NullString)
		case searchkeyword.ForeignKeys[0]: // user_searched_keywords
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type SearchKeyword", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SearchKeyword fields.
func (sk *SearchKeyword) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case searchkeyword.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			sk.ID = int(value.Int64)
		case searchkeyword.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				sk.Title = value.String
			}
		case searchkeyword.FieldRate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field rate", values[i])
			} else if value.Valid {
				sk.Rate = uint16(value.Int64)
			}
		case searchkeyword.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_searched_keywords", value)
			} else if value.Valid {
				sk.user_searched_keywords = new(int)
				*sk.user_searched_keywords = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the SearchKeyword entity.
func (sk *SearchKeyword) QueryUser() *UserQuery {
	return (&SearchKeywordClient{config: sk.config}).QueryUser(sk)
}

// Update returns a builder for updating this SearchKeyword.
// Note that you need to call SearchKeyword.Unwrap() before calling this method if this SearchKeyword
// was returned from a transaction, and the transaction was committed or rolled back.
func (sk *SearchKeyword) Update() *SearchKeywordUpdateOne {
	return (&SearchKeywordClient{config: sk.config}).UpdateOne(sk)
}

// Unwrap unwraps the SearchKeyword entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (sk *SearchKeyword) Unwrap() *SearchKeyword {
	tx, ok := sk.config.driver.(*txDriver)
	if !ok {
		panic("ent: SearchKeyword is not a transactional entity")
	}
	sk.config.driver = tx.drv
	return sk
}

// String implements the fmt.Stringer.
func (sk *SearchKeyword) String() string {
	var builder strings.Builder
	builder.WriteString("SearchKeyword(")
	builder.WriteString(fmt.Sprintf("id=%v", sk.ID))
	builder.WriteString(", title=")
	builder.WriteString(sk.Title)
	builder.WriteString(", rate=")
	builder.WriteString(fmt.Sprintf("%v", sk.Rate))
	builder.WriteByte(')')
	return builder.String()
}

// SearchKeywords is a parsable slice of SearchKeyword.
type SearchKeywords []*SearchKeyword

func (sk SearchKeywords) config(cfg config) {
	for _i := range sk {
		sk[_i].config = cfg
	}
}
