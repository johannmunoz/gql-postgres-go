package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
)

func (r *entityResolver) FindManufacturerByID(ctx context.Context, id string) (*ent.Manufacturer, error) {
	panic(fmt.Errorf("FindManufacturerByID not implemented"))
}

func (r *entityResolver) FindProductByID(ctx context.Context, id string) (*ent.Product, error) {
	productId, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Product.Query().Where(product.IDEQ(productId)).Only(ctx)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
