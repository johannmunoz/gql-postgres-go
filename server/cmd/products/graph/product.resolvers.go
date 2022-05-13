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

func (r *mutationResolver) ProductCreate(ctx context.Context, input *model.NewProduct) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) ID(ctx context.Context, obj *ent.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context) ([]*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Product(ctx context.Context, id string) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
