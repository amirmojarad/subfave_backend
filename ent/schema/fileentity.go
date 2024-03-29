package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// File holds the schema definition for the File entity.
type FileEntity struct {
	ent.Schema
}

// Fields of the File.
func (FileEntity) Fields() []ent.Field {
	return []ent.Field{
		field.String("path").NotEmpty().Annotations(entproto.Field(2)),
		field.String("name").NotEmpty().Annotations(entproto.Field(3)),
		field.Int16("size").Annotations(entproto.Field(4)),
		field.Bool("deleted").Annotations(entproto.Field(5)),
		field.Time("created_date").Annotations(entproto.Field(6)),
	}
}

// Edges of the File.
func (FileEntity) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("files").Unique().Annotations(entproto.Field(7), entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("words", Word.Type).Annotations(entproto.Field(8), entsql.Annotation{
			OnDelete: entsql.Cascade,
		}),
	}
}

func (FileEntity) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
