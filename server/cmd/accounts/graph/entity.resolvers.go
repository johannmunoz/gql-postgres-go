package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/user"
)

func (r *entityResolver) FindUserByID(ctx context.Context, id string) (*ent.User, error) {
	userId, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.User.Query().Where(user.IDEQ(userId)).Only(ctx)
}

// Entity returns generated.EntityResolver implementation.
func (r *Resolver) Entity() generated.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
