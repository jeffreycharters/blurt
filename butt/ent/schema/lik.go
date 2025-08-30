package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Lik holds the schema definition for the Lik entity.
type Lik struct {
	ent.Schema
}

func (Lik) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "blurt_id"),
	}
}

// Fields of the Lik.
func (Lik) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_id", uuid.UUID{}),
		field.UUID("blurt_id", uuid.UUID{}),
	}
}

// Edges of the Lik.
func (Lik) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("blurt", Blurt.Type).
			Unique().
			Required().
			Field("blurt_id"),
	}
}
