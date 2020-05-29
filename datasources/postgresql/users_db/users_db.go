package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	postgresql_user_username = "username"
	postgresql_user_password = "password"
	postgresql_user_host     = "host"
	postgresql_user_port     = "port"
	postgresql_user_database = "database"
)

var (
	Client *sql.DB

	username = os.Getenv(postgresql_user_username)
	password = os.Getenv(postgresql_user_password)
	host     = os.Getenv(postgresql_user_host)
	port     = os.Getenv(postgresql_user_port)
	database = os.Getenv(postgresql_user_database)
)

func init() {
	var err error

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		database,
	)

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connected")
}
