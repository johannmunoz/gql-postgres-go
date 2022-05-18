package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/johannmunoz/gql_postgres_go/ent"
)

func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		log.Fatal(err)
	}
	return r.client.Noder(ctx, uuid)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	var UUIDs []uuid.UUID
	for _, x := range ids {
		uuid, err := uuid.Parse(x)
		if err != nil {
			log.Fatal(err)
		}
		UUIDs = append(UUIDs, uuid)
	}
	return r.client.Noders(ctx, UUIDs)
}
