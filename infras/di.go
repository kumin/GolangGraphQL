package infras

import "github.com/google/wire"

var GraphSet = wire.NewSet(
	GetMysqlCfgs,
	BuildConnection,
)
