package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/products/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/model"
)

func (r *manufacturerResolver) ID(ctx context.Context, obj *ent.Manufacturer) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ManufacturerCreate(ctx context.Context, input *model.NewManufaturer) (*ent.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Manufacturer(ctx context.Context, obj *ent.Product) (*ent.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Manufacturers(ctx context.Context) ([]*ent.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Manufacturer(ctx context.Context, id string) (*ent.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
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
