package config

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"golang-echo-graphql-example/graph"
	"golang-echo-graphql-example/graph/generated"
)

func NewGraphQLServer(resolver *graph.Resolver) *handler.Server {
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	return srv
}
