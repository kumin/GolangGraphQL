// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/google/wire"
	"github.com/kumin/GolangGraphQL/handlers/v1/http"
	"github.com/kumin/GolangGraphQL/infras"
	"github.com/kumin/GolangGraphQL/repos/bind"
	"github.com/kumin/GolangGraphQL/repos/mysql"
)

// Injectors from wire.go:

func BuildServer() (*HttpServer, error) {
	mysqlConfiguration := infras.GetMysqlCfgs()
	db, err := infras.BuildConnection(mysqlConfiguration)
	if err != nil {
		return nil, err
	}
	productMysqlRepo := mysql.NewProductMysqlRepo(db)
	productCtlHandler := http.NewProductCtlHandler(productMysqlRepo)
	httpServer := NewServer(productMysqlRepo, productCtlHandler)
	return httpServer, nil
}

// wire.go:

var GraphSet = wire.NewSet(bind.GraphSet, http.GraphSet, infras.GraphSet, NewServer)
