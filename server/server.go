package main

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gqlgorm/generated"
	"gqlgorm/resolvers"
	db "gqlgorm/utils/database"
	"gqlgorm/utils/middleware"
)
func graphqlHandler() gin.HandlerFunc  {
	h := handler.GraphQL(generated.NewExecutableSchema(generated.Config{Resolvers:&resolvers.Resolver{}}))
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer,c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc  {
	h := handler.Playground("GraphQL","/query")
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer,c.Request)
	}
}
func main() {
    db.GetInstance()
    defer db.Close()
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(middleware.GinContextToContextMiddleware())
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:  []string{"content-type"},
		ExposeHeaders: []string{"X-Total-Count"},
	}))
	r.POST("/query",graphqlHandler())
	r.GET("/",playgroundHandler())
	r.Run("localhost:8080")

}
