package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/generated"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	panic(fmt.Errorf("FindUserByID not implemented"))
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
