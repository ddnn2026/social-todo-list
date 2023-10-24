package transport

import (
	"net/http"

	"social-todo-list/common"
	"social-todo-list/modules/item/business"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/storage"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemCreate
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		store := storage.NewSQLStore(db)
		businessStore := business.NewCreateItem(store)

		if err := businessStore.CreateNewItem(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// result := db.Table("todo_items").Create(&data)
		// fmt.Println(result)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}
