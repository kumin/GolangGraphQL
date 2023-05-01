package infras

import (
	"database/sql"

	"github.com/kumin/GolangGraphQL/helpers/envx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfiguration struct {
	URI               string
	MaxConnection     int
	MaxIdleConnection int
}

func BuildConnection(configs *MysqlConfiguration) (*gorm.DB, error) {
	conn, err := sql.Open("mysql", configs.URI)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: conn,
	}), &gorm.Config{
		QueryFields: true,
	})
	if err != nil {
		return nil, err
	}
	pool, err := db.DB()
	if err != nil {
		return nil, err
	}
	pool.SetMaxOpenConns(configs.MaxConnection)
	pool.SetMaxIdleConns(configs.MaxIdleConnection)

	return db, nil
}

func GetMysqlCfgs() *MysqlConfiguration {
	URI := envx.GetString("MYSQL_ADDRS", "root:root@tcp(localhost:3306)/kumin_store?charset=utf8&parseTime=True&loc=Local&multistatement=true")
	MaxConns := envx.GetInt("MAX_CONNECTIONS", 10)
	MaxIdleConnection := envx.GetInt("MAX_IDLECONNECTIONS", 10)

	return &MysqlConfiguration{
		URI:               URI,
		MaxConnection:     MaxConns,
		MaxIdleConnection: MaxIdleConnection,
	}
}
