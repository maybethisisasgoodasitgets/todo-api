package app

import "todo/app/controllers/ping"




func mapUrls(){
	router.GET("/ping",ping.Ping)
	router.POST("/item",item.Create)
}