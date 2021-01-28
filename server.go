package main

import (
	"github.com/maapteh/graphql-golang/generated"
	"github.com/maapteh/graphql-golang/resolver"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

// Defining the Health handler which returns simple OK
func healthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Status(http.StatusOK)
	}
}

// Defining the Graphql handler
func graphqlHandler() gin.HandlerFunc {

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}})

	h := handler.NewDefaultServer(schema)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/graphql")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {

	ginMode := os.Getenv("GIN_MODE")

	// Setting up Gin
    // TODO: Disable log's color ? on production for example we dont need fancy logging
	// gin.DisableConsoleColor()

	r := gin.New()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release. By default gin.DefaultWriter = os.Stdout
	if ginMode != "release" {
		r.Use(gin.Logger())
	}
	
	// locally and pod itself
	r.GET("/", playgroundHandler())
	r.POST("/graphql", graphqlHandler())

	// health check
	r.GET("/.well-known/server-health", healthHandler())

	r.Run()
}
