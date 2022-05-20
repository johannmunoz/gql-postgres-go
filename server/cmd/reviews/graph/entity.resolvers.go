package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
	"github.com/johannmunoz/gql_postgres_go/ent/review"
	"github.com/johannmunoz/gql_postgres_go/ent/user"
)

func (r *entityResolver) FindProductByID(ctx context.Context, id uuid.UUID) (*ent.Product, error) {
	return r.client.Product.Query().Where(product.IDEQ(id)).Only(ctx)
}

func (r *entityResolver) FindReviewByID(ctx context.Context, id string) (*ent.Review, error) {
	reviewId, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Review.Query().Where(review.IDEQ(reviewId)).Only(ctx)
}

func (r *entityResolver) FindUserByID(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	return r.client.User.Query().Where(user.IDEQ(id)).Only(ctx)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
