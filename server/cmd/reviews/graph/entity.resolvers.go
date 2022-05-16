package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/generated"
)

func (r *entityResolver) FindProductByID(ctx context.Context, id string) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindReviewByID(ctx context.Context, id string) (*ent.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
