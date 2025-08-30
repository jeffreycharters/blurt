package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

// Blurt holds the schema definition for the Blurt entity.
type Blurt struct {
	ent.Schema
}

// Fields of the Blurt.
func (Blurt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("content").
			NotEmpty().
			MaxLen(14),
	}
}

// Edges of the Blurt.
func (Blurt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("likd_users", User.Type).Ref("likd_blurts").Through("liks", Lik.Type),
		edge.From("author", User.Type).
			Unique().
			Ref("blurts"),
	}
}

func (Blurt) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
