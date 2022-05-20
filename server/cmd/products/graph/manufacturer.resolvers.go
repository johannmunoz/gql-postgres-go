package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/model"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
)

func (r *mutationResolver) ManufacturerCreate(ctx context.Context, input model.NewManufacturer) (*ent.Manufacturer, error) {
	return r.client.Manufacturer.Create().SetName(input.Name).Save(ctx)
}

func (r *queryResolver) Manufacturers(ctx context.Context, before *ent.Cursor, after *ent.Cursor, first *int, last *int, orderBy *ent.ManufacturerOrder, where *ent.ManufacturerWhereInput) (*ent.ManufacturerConnection, error) {
	return r.client.Manufacturer.Query().Paginate(ctx, after, first, before, last, ent.WithManufacturerOrder(orderBy), ent.WithManufacturerFilter(where.Filter))
}

func (r *queryResolver) Manufacturer(ctx context.Context, id string) (*ent.Manufacturer, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Manufacturer.Query().Where(manufacturer.IDEQ(uuid)).Only(ctx)
}

// Manufacturer returns generated.ManufacturerResolver implementation.
func (r *Resolver) Manufacturer() generated.ManufacturerResolver { return &manufacturerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type manufacturerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
