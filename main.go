package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"social-todo-list/common"
	"social-todo-list/modules/item/model"
	"social-todo-list/modules/item/transport"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:mypass@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("", getAll(db))
		v1.POST("", transport.CreateItem(db))
		v1.GET("/:id", getItem(db))
		v1.PATCH("/:id", updateItem(db))
		v1.DELETE("/:id", deleteItem(db))
	}

	r.GET("/ping", getPing)
	r.Run(":8088")
}

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getAll(db *gorm.DB) func(c *gin.Context) {
	var result []model.TodoItem
	return func(c *gin.Context) {
		if err := db.Table("todo_items").Find(&result).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}

func getItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItem
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		if err := db.Table("todo_items").Where("id = ?", id).First(&data).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}

func updateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data model.TodoItemUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		if err := db.Table("todo_items").Where("id = ?", id).Updates(&data).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println(data)
		c.JSON(http.StatusOK, gin.H{
			"data": true,
		})
	}
}

func deleteItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
		}

		if err := db.Table("todo_items").Where("id = ?", id).Delete(nil).Error; err != nil {
			c.JSON(http.StatusBadGateway, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
