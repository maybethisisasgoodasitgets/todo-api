package tododb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

const (
	postgresTodoUsername = "potgres"
	postgresTodoPassword = "admin"
	postgresTodoHost     = "locolhost"
	postgresTodoSchema   = "reza"
)

var (
	Client *sql.DB

	username = os.Getenv(postgresTodoUsername)
	password = os.Getenv(postgresTodoPassword)
	host     = os.Getenv(postgresTodoHost)
	schema   = os.Getenv(postgresTodoSchema)
)

func init(){

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",username,password,host,schema)

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
