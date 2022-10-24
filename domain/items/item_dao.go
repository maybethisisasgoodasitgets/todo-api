package items

import (
	"errors"
	tododb "todo/datasource/postgres/todo_db"

	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)
const (
	queryInsertItem = "INSERT INTO list (item,done) VALUES(?,?);"
)



func (item *Item)Save() rest_errors.RestErr{
stmt, err := tododb.Client.Prepare(queryInsertItem)

if err != nil{
	logger.Error("error when trying to prepare the insert item statement",err)
	return rest_errors.NewInternalServerError("error when trying to insert item",errors.New("database error"))

	//... to be continued
}

}