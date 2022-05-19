package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("name").Annotations(
			entgql.OrderField("name"),
		),
		field.String("upc").Annotations(
			entgql.OrderField("upc"),
		),
		field.Int("price").Default(0).Annotations(
			entgql.OrderField("price"),
		),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("manufacturer", Manufacturer.Type).
			Ref("products").
			Unique(),
		edge.To("reviews", Review.Type),
	}
}
