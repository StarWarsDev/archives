package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/StarWarsDev/archives/internal/gql"
	"github.com/StarWarsDev/archives/internal/gql/resolvers"
	"log"
	"net/http"
	"os"
)

const defaultPort = "3001"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	http.Handle("/graphql", handler.GraphQL(gql.NewExecutableSchema(gql.Config{Resolvers: &resolvers.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
