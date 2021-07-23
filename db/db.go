package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var _db *sql.DB

func Initialize() {
	hostname := viper.GetString("mariadb.address")
	user := viper.GetString("mariadb.username")
	pass := viper.GetString("mariadb.password")
	port := viper.GetString("mariadb.port")
	database := viper.GetString("mariadb.database")

	// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	var connectionStr = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, hostname, port, database)
	_db, _ = sql.Open("mysql", connectionStr)
}

func GetDB() *sql.DB {
	if _db == nil {
		Initialize()
	}

	return _db
}

func Ping() error {
	return _db.Ping()
}
