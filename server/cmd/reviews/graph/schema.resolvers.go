package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/model"
)

func (r *productResolver) ID(ctx context.Context, obj *ent.Product) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Manufacturer(ctx context.Context, obj *ent.Product) (*model.Manufacturer, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *productResolver) Reviews(ctx context.Context, obj *ent.Product) ([]*ent.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *reviewResolver) ID(ctx context.Context, obj *ent.Review) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *reviewResolver) Author(ctx context.Context, obj *ent.Review) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *reviewResolver) Product(ctx context.Context, obj *ent.Review) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

type productResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
