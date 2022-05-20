package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
)

func (r *entityResolver) FindManufacturerByID(ctx context.Context, id uuid.UUID) (*ent.Manufacturer, error) {
	return r.client.Manufacturer.Query().Where(manufacturer.IDEQ(id)).Only(ctx)
}

func (r *entityResolver) FindProductByID(ctx context.Context, id uuid.UUID) (*ent.Product, error) {
	return r.client.Product.Query().Where(product.IDEQ(id)).Only(ctx)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
