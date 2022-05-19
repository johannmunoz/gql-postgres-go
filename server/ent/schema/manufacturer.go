package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Manufacturer holds the schema definition for the Manufacturer entity.
type Manufacturer struct {
	ent.Schema
}

// Fields of the Manufacturer.
func (Manufacturer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").Annotations(
			entgql.OrderField("name"),
		),
	}
}

// Edges of the Manufacturer.
func (Manufacturer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("products", Product.Type),
	}
}
