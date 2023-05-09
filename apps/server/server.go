package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/kumin/GolangGraphQL/graph"
	"github.com/kumin/GolangGraphQL/handlers/v1/http"
	"github.com/kumin/GolangGraphQL/infras"
	"github.com/kumin/GolangGraphQL/repos/mysql"
	"golang.org/x/sync/errgroup"
)

func graphQLHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{
			Resolvers: &graph.Resolver{},
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

func main() {
	mysqlConfig := infras.GetMysqlCfgs()
	mysqlConn, err := infras.BuildConnection(mysqlConfig)
	if err != nil {
		panic(err)
	}
	productRepo := mysql.NewProductMysqlRepo(mysqlConn)
	productHandler := http.NewProductCtlHandler(productRepo)
	router := gin.Default()
	router.POST("/v1/product/create", productHandler.CreateProduct)
	router.GET("/v1/product/listing", productHandler.ListProducts)
	router.GET("/", playgroundHandler())
	router.POST("/v1/query", graphQLHandler())

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()
	eg, _ := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return router.Run("localhost:8080")
	})
	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
