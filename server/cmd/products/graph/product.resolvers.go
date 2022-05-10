package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/model"
)

func (r *queryResolver) TopProducts(ctx context.Context, first *int) ([]*model.Product, error) {
	return hats, nil
}
