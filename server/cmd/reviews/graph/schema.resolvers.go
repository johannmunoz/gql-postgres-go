package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/reviews/graph/model"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
	"github.com/johannmunoz/gql_postgres_go/ent/user"
)

func (r *manufacturerResolver) ID(ctx context.Context, obj *ent.Manufacturer) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ReviewCreate(ctx context.Context, input model.NewReview, authorID string, productID string) (*ent.Review, error) {
	authorUUID, err := uuid.Parse(authorID)
	if err != nil {
		log.Fatal(err)
	}
	author, err := r.client.User.Query().Where(user.IDEQ(authorUUID)).Only(ctx)
	if err != nil {
		log.Fatal(err)
	}
	productUUID, err := uuid.Parse(productID)
	if err != nil {
		log.Fatal(err)
	}
	product, err := r.client.Product.Query().Where(product.IDEQ(productUUID)).Only(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return r.client.Review.Create().SetBody(input.Body).SetProductID(product.ID).SetAuthorID(author.ID).Save(ctx)
}

func (r *productResolver) ID(ctx context.Context, obj *ent.Product) (string, error) {
	return obj.ID.String(), nil
}

func (r *productResolver) Manufacturer(ctx context.Context, obj *ent.Product) (*ent.Manufacturer, error) {
	panic(fmt.Errorf("not implemented Manufacturer"))
}

func (r *productResolver) Reviews(ctx context.Context, obj *ent.Product) ([]*ent.Review, error) {
	panic(fmt.Errorf("not implemented Reviews"))
}

func (r *queryResolver) Reviews(ctx context.Context) ([]*ent.Review, error) {
	return r.client.Review.Query().All(ctx)
}

func (r *queryResolver) Review(ctx context.Context, id string) (*ent.Review, error) {
	panic(fmt.Errorf("not implemented Review"))
}

func (r *reviewResolver) ID(ctx context.Context, obj *ent.Review) (string, error) {
	return obj.ID.String(), nil
}

func (r *reviewResolver) Author(ctx context.Context, obj *ent.Review) (*ent.User, error) {
	authorId, err := obj.QueryAuthor().OnlyID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.User.Query().Where(user.IDEQ(authorId)).Only(ctx)
}

func (r *reviewResolver) Product(ctx context.Context, obj *ent.Review) (*ent.Product, error) {
	panic(fmt.Errorf("not implemented Product"))
}

func (r *userResolver) ID(ctx context.Context, obj *ent.User) (string, error) {
	return obj.ID.String(), nil
}

func (r *userResolver) Reviews(ctx context.Context, obj *ent.User) ([]*ent.Review, error) {
	panic(fmt.Errorf("not implemented Reviews"))
}

// Manufacturer returns generated.ManufacturerResolver implementation.
func (r *Resolver) Manufacturer() generated.ManufacturerResolver { return &manufacturerResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Review returns generated.ReviewResolver implementation.
func (r *Resolver) Review() generated.ReviewResolver { return &reviewResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type manufacturerResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reviewResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *userResolver) Email(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented Email"))
}
func (r *userResolver) Username(ctx context.Context, obj *ent.User) (string, error) {
	panic(fmt.Errorf("not implemented Username"))
}
