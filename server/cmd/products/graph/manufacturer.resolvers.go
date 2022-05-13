package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/model"
)

func (r *manufacturerResolver) ID(ctx context.Context, obj *ent.Manufacturer) (string, error) {
	return obj.ID.String(), nil
}

func (r *mutationResolver) ManufacturerCreate(ctx context.Context, input model.NewManufaturer) (*ent.Manufacturer, error) {
	return r.client.Manufacturer.Create().SetName(input.Name).Save(ctx)
}

func (r *productResolver) Manufacturer(ctx context.Context, obj *ent.Product) (*ent.Manufacturer, error) {
	manufacturerId, err := obj.QueryManufacturer().OnlyID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Manufacturer.Query().Where(manufacturer.IDEQ(manufacturerId)).Only(ctx)
}

func (r *queryResolver) Manufacturers(ctx context.Context) ([]*ent.Manufacturer, error) {
	return r.client.Manufacturer.Query().All(ctx)
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
