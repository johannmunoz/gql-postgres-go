package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/model"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input model.NewUser) (*ent.User, error) {
	return r.client.User.Create().SetEmail(input.Email).SetUsername(input.Username).Save(ctx)
	// if err != nil {
	// 	log.Fatalf("failed creating a todo: %v", err)
	// }
	// return user, nil
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
