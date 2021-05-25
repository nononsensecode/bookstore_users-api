package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host     = "mysql_users_host"
	mysql_users_port     = "mysql_users_port"
	mysql_users_schema   = "mysql_users_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	port     = os.Getenv(mysql_users_port)
	schema   = os.Getenv(mysql_users_schema)
)

func init() {

	datasource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, password, host, port, schema,
	) //"kaushik", "redhat", "127.0.0.1", "users_db",

	var err error
	Client, err = sql.Open("mysql", datasource)
	if err != nil {
		panic(err)
	}

	Ping()

	log.Println("database successfully configured")
}

func Ping() {
	if err := Client.Ping(); err != nil {
		panic(err)
	}
}
