package db

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"user/api/internal/config"
)

func NewMysql(mysqlConfig config.MysqlConfig) sqlx.SqlConn {
	mysql := sqlx.NewMysql(mysqlConfig.DataSource)
	return mysql
}
