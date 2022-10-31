package app

import (
	"todo/controllers/ping"
	"todo/controllers/items"
)




func mapUrls(){
	router.GET("/ping",ping.Ping)
	router.GET("/item/create/:item",items.Create)
}