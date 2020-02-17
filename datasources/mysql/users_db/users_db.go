package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

const (
	MYSQL_USERNAME     = "MYSQL_USERNAME"
	MYSQL_PASSWORD     = "MYSQL_PASSWORD"
	MYSQL_HOST         = "MYSQL_HOST"
	MYSQL_USERS_SCHEMA = "MYSQL_USERS_SCHEMA"
)

var (
	ClientDB *sql.DB

	username = os.Getenv(MYSQL_USERNAME)
	password = os.Getenv(MYSQL_PASSWORD)
	host     = os.Getenv(MYSQL_HOST)
	schema   = os.Getenv(MYSQL_USERS_SCHEMA)
)

func init() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		username, // User
		password, // Password
		host,     // host:port
		schema)   //  DB schema name

	var connectionErr error
	ClientDB, connectionErr = sql.Open("mysql", dataSource)
	if connectionErr != nil {
		panic(connectionErr)
	}

	if err := ClientDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Database connected")

}
