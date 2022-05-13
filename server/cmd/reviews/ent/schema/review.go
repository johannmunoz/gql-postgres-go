package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("body"),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("product", Product.Type).
			Ref("reviews").
			Unique(),
	}
}
