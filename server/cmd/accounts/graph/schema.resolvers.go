package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/model"
	"github.com/johannmunoz/gql_postgres_go/ent"
)

func (r *mutationResolver) UserCreate(ctx context.Context, input model.NewUser) (*ent.User, error) {
	return r.client.User.Create().SetEmail(input.Email).SetUsername(input.Username).Save(ctx)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("Me not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, before *ent.Cursor, after *ent.Cursor, first *int, last *int, orderBy *ent.UserOrder) (*ent.UserConnection, error) {
	return r.client.User.Query().Paginate(ctx, after, first, before, last, ent.WithUserOrder(orderBy))
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
