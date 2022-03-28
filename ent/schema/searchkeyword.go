package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// SearchKeyword holds the schema definition for the SearchKeyword entity.
type SearchKeyword struct {
	ent.Schema
}

// Fields of the SearchKeyword.
func (SearchKeyword) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Unique(),
		field.Uint16("rate").Default(0),
	}
}

// Edges of the SearchKeyword.
func (SearchKeyword) Edges() []ent.Edge {
	return nil
}