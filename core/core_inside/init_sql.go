package core_inside

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func DbSqlActuator() {
	db, err := sql.Open("mysql", MysqlConfig.Datasource)
	if err != nil {
		fmt.Println("链接数据库失败：", err.Error())
	}
	Db = db
}
