package tododb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/lib/pq"
)

const (
	postgresTodoUsername = "postgres_todo_username"
	postgresTodoPassword = "postgres_todo_password"
	postgresTodoHost     = "postgres_todo_host"
	postgresTodoSchema   = "postgres_todo_schema"
)

var (
	Client *sql.DB

	username = os.Getenv(postgresTodoUsername)
	password = os.Getenv(postgresTodoPassword)
	host     = os.Getenv(postgresTodoHost)
	schema   = os.Getenv(postgresTodoSchema)
)

func init(){

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",username,password,host,schema)
    fmt.Println(dataSourceName)
	var err error
	Client, err = sql.Open("postgres",dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)

	}

	log.Println("database was connected successfully!")


}
