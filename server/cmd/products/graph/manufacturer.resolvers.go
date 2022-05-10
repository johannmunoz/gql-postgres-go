package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/model"
)

func (r *mutationResolver) ManufacturerCreate(ctx context.Context, input *model.NewManufaturer) ([]*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Manufacturers(ctx context.Context) ([]*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Manufacturer(ctx context.Context, id string) (*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
