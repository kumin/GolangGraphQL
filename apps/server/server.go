package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kumin/GolangGraphQL/handlers/v1/http"
	"github.com/kumin/GolangGraphQL/infras"
	"github.com/kumin/GolangGraphQL/repos/mysql"
)

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
  router.Run("localhost:8080")
}
