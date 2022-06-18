// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"interface_project/ent/category"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Category is the model entity for the Category schema.
type Category struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// CreatedDate holds the value of the "created_date" field.
	CreatedDate time.Time `json:"created_date,omitempty"`
	// UpdatedDate holds the value of the "updated_date" field.
	UpdatedDate time.Time `json:"updated_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CategoryQuery when eager-loading is set.
	Edges CategoryEdges `json:"edges"`
}

// CategoryEdges holds the relations/edges for other nodes in the graph.
type CategoryEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// CategoryWords holds the value of the category_words edge.
	CategoryWords []*Word `json:"category_words,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// CategoryWordsOrErr returns the CategoryWords value or an error if the edge
// was not loaded in eager-loading.
func (e CategoryEdges) CategoryWordsOrErr() ([]*Word, error) {
	if e.loadedTypes[1] {
		return e.CategoryWords, nil
	}
	return nil, &NotLoadedError{edge: "category_words"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Category) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			values[i] = new(sql.NullInt64)
		case category.FieldTitle:
			values[i] = new(sql.NullString)
		case category.FieldCreatedDate, category.FieldUpdatedDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Category", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Category fields.
func (c *Category) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case category.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case category.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case category.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_date", values[i])
			} else if value.Valid {
				c.CreatedDate = value.Time
			}
		case category.FieldUpdatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_date", values[i])
			} else if value.Valid {
				c.UpdatedDate = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Category entity.
func (c *Category) QueryUser() *UserQuery {
	return (&CategoryClient{config: c.config}).QueryUser(c)
}

// QueryCategoryWords queries the "category_words" edge of the Category entity.
func (c *Category) QueryCategoryWords() *WordQuery {
	return (&CategoryClient{config: c.config}).QueryCategoryWords(c)
}

// Update returns a builder for updating this Category.
// Note that you need to call Category.Unwrap() before calling this method if this Category
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Category) Update() *CategoryUpdateOne {
	return (&CategoryClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Category entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Category) Unwrap() *Category {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Category is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Category) String() string {
	var builder strings.Builder
	builder.WriteString("Category(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", title=")
	builder.WriteString(c.Title)
	builder.WriteString(", created_date=")
	builder.WriteString(c.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", updated_date=")
	builder.WriteString(c.UpdatedDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Categories is a parsable slice of Category.
type Categories []*Category

func (c Categories) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
