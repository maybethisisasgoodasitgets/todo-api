package items

import (
	"net/http"
	"todo/domain/items"
	"todo/services"

	"github.com/gin-gonic/gin"
)


func Create(c *gin.Context){
	var item items.Item
	item.Item = c.Param("item")

	result, saveErr := services.ItemsService.CreateItem(item)
	if saveErr != nil {
		c.JSON(saveErr.Status(),saveErr)
		return
	}

	c.JSON(http.StatusCreated,result.Marshall(c.Header("X-Public")== "true"))


 
}