package main

import (
	"log"
	"net/http"
	"os"

	resolver "bitbucket.org/antuitinc/esp-cp-api/internal/graph/cp-api"

	"bitbucket.org/antuitinc/esp-cp-api/internal/access"
	"bitbucket.org/antuitinc/esp-cp-api/internal/db"
	"bitbucket.org/antuitinc/esp-cp-api/internal/graph/cp-api/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := resolver.Resolver{}
	if err := r.LoadTemplates(); err != nil {
		log.Fatal(err)
	}
	r.DBClient = db.New()

	authMiddleware := access.AuthMiddleware()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &r}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", authMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
