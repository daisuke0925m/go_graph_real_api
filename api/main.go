package main

import (
	"log"
	"net/http"
	"time"

	"api/graph/generated"
	"api/src/config"
	"api/src/infra/db"
	"api/src/infra/redis"
	"api/src/resolver"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// Defining the Graphql handlerhandler/transport
func graphqlHandler() gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	resolver := resolver.New()
	h := handler.GraphQL(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: resolver,
			},
		),
		handler.WebsocketKeepAliveDuration(10*time.Second),
		handler.WebsocketUpgrader(websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			HandshakeTimeout: 10 * time.Second,
		}),
	)

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	config.Load()
	_ = db.NewDb()
	_ = redis.New()

	r := gin.Default()
	h := graphqlHandler()

	r.POST("/query", h)
	r.GET("/query", h)
	r.GET("/", playgroundHandler())
	r.Run(":" + config.Conf.App.Port)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Conf.App.Port)
}
