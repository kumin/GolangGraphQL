package bind

import (
	"github.com/google/wire"
	"github.com/kumin/GolangGraphQL/repos"
	"github.com/kumin/GolangGraphQL/repos/mysql"
)

var GraphSet = wire.NewSet(
	mysql.NewProductMysqlRepo,
	wire.Bind(new(repos.ProductRepo), new(*mysql.ProductMysqlRepo)),
)
