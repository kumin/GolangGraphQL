//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"github.com/kumin/GolangGraphQL/handlers/v1/http"
	"github.com/kumin/GolangGraphQL/infras"
	"github.com/kumin/GolangGraphQL/repos/bind"
)

var GraphSet = wire.NewSet(
	bind.GraphSet,
	http.GraphSet,
	infras.GraphSet,
	NewServer,
)

func BuildServer() (*HttpServer, error) {
	wire.Build(
		GraphSet,
	)
	return nil, nil
}
