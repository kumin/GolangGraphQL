package server

import (
	"fmt"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/kumin/GolangGraphQL/graph"
	"github.com/kumin/GolangGraphQL/handlers/v1/http"
	"github.com/kumin/GolangGraphQL/helpers/envx"
	"github.com/kumin/GolangGraphQL/repos"
	"golang.org/x/net/context"
)

type HttpServer struct {
	router *gin.Engine
}

func NewServer(
	productRepo repos.ProductRepo,
	productHandler *http.ProductCtlHandler,
) *HttpServer {
	router := gin.Default()
	router.POST("/v1/product/create", productHandler.CreateProduct)
	router.GET("/v1/product/listing", productHandler.ListProducts)
	router.GET("/", playgroundHandler())
	router.POST("/v1/query", graphQLHandler(productRepo))
	router.POST("/v1/mutation", graphQLHandler(productRepo))

	return &HttpServer{
		router: router,
	}
}

func (s *HttpServer) Start(ctx context.Context) error {
	port := envx.GetInt("SERVING_PORT", 8080)
	return s.router.Run(fmt.Sprintf(":%d", port))
}

func graphQLHandler(producRepo repos.ProductRepo) gin.HandlerFunc {
	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{ProductRepo: producRepo},
		}),
	)

	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
