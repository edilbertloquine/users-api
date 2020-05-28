package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

const (
	postgresql_user_username = "username"
	postgresql_user_password = "password"
	postgresql_user_host     = "host"
	postgresql_user_port     = "port"
	postgresql_user_schema   = "schema"
)

var (
	Client *sql.DB

	username = os.Getenv(postgresql_user_username)
	password = os.Getenv(postgresql_user_password)
	host     = os.Getenv(postgresql_user_host)
	port     = os.Getenv(postgresql_user_port)
	schema   = os.Getenv(postgresql_user_schema)
)

func init() {
	var err error
	var portnum int64

	portnum, err = strconv.ParseInt(port, 10, 64)
	if err != nil {
		panic(err)
	}

	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		portnum,
		username,
		password,
		schema,
	)

	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database connected")
}
