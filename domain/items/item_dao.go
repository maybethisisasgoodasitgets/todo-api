package items

import (
	"errors"
	tododb "todo/datasource/postgres/todo_db"

	"github.com/federicoleon/bookstore_utils-go/logger"
	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)
const (
	queryInsertItem = "INSERT INTO list (item,done) VALUES($1,$2);"
	notDone = false
)



func (item *Item) Save() rest_errors.RestErr{
stmt, err := tododb.Client.Prepare(queryInsertItem)

if err != nil{
	logger.Error("error when trying to prepare the insert item statement",err)
	return rest_errors.NewInternalServerError("error when trying to insert item",errors.New("database error"))

	
}
defer stmt.Close()
item.Done = notDone
_,saveErr := stmt.Exec(item.Item,item.Done)
if saveErr != nil{
	logger.Error("error when trying to save item to the database!",saveErr)

	return rest_errors.NewInternalServerError("error when trying to save item to database!",errors.New("database error"))
}

return nil

}