// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/ent"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph/generated"
)

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
