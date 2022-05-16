package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/johannmunoz/gql_postgres_go/cmd/accounts/graph"
	"github.com/johannmunoz/gql_postgres_go/ent"
	_ "github.com/lib/pq"
)

const defaultPort = "4001"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secret"
	dbname   = "testing_db"
)

func main() {

	var entOptions []ent.Option
	entOptions = append(entOptions, ent.Debug())

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	client, err := ent.Open("postgres", psqlInfo, entOptions...)

	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); !errors.Is(err, nil) {
		log.Fatalf("Error: failed creating schema resources %v\n", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewSchema(client))
	srv.Use(&debug.Tracer{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
