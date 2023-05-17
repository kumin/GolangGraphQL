package http

import "github.com/google/wire"

var GraphSet = wire.NewSet(
	NewProductCtlHandler,
)
