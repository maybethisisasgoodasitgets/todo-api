package services

import (
	"todo/domain/items"

	"github.com/federicoleon/bookstore_utils-go/rest_errors"
)


var (
	ItemsService itemsServiceInterface = &itemsService{}
)
type itemsService struct{}
type itemsServiceInterface interface {
	CreateItem(items.Item) (*items.Item, rest_errors.RestErr)
}

func (s *itemsService) CreateItem(item items.Item) (*items.Item, rest_errors.RestErr) {
if err:= item.Save(); err != nil {
	return nil,err
}
return &item,nil
}
