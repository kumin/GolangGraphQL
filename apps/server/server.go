package main

import (
	"github.com/kumin/GolangGraphQL/handler/v1/http"
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
	productHandler := http.NewProductHandler(productRepo)
}
