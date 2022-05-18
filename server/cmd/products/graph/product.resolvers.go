package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/generated"
	"github.com/johannmunoz/gql_postgres_go/cmd/products/graph/model"
	"github.com/johannmunoz/gql_postgres_go/ent"
	"github.com/johannmunoz/gql_postgres_go/ent/manufacturer"
	"github.com/johannmunoz/gql_postgres_go/ent/product"
)

func (r *manufacturerResolver) Products(ctx context.Context, obj *ent.Manufacturer, before *ent.Cursor, after *ent.Cursor, first *int, last *int) (*ent.ProductConnection, error) {
	return r.client.Product.Query().Where(product.HasManufacturerWith(manufacturer.IDEQ(obj.ID))).Paginate(ctx, after, first, before, last)
}

func (r *mutationResolver) ProductCreate(ctx context.Context, input model.NewProduct, manufacturerID string) (*ent.Product, error) {
	manufacturerUUID, err := uuid.Parse(manufacturerID)
	if err != nil {
		log.Fatal(err)
	}
	manufacturer, err := r.client.Manufacturer.Query().Where(manufacturer.IDEQ(manufacturerUUID)).Only(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Product.Create().SetName(input.Name).SetUpc(input.Upc).SetPrice(input.Price).SetManufacturerID(manufacturer.ID).Save(ctx)
}

func (r *productResolver) ID(ctx context.Context, obj *ent.Product) (string, error) {
	return obj.ID.String(), nil
}

func (r *queryResolver) Products(ctx context.Context, before *ent.Cursor, after *ent.Cursor, first *int, last *int) (*ent.ProductConnection, error) {
	return r.client.Product.Query().Paginate(ctx, after, first, before, last)
}

func (r *queryResolver) Product(ctx context.Context, id string) (*ent.Product, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Product.Query().Where(product.IDEQ(uuid)).Only(ctx)
}

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

type productResolver struct{ *Resolver }
