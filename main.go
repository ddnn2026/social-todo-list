package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"social-todo-list/common"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	common.SQLmodel
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"id;"`
	Title       string `json:"title" gorm:"title;"`
	Description string `json:"description" gorm:"description;"`
	Status      string `json:"status" gorm:"status;"`
}

type TodoItemUpdate struct {
	Title       string `json:"title" gorm:"title;"`
	Description string `json:"description" gorm:"description;"`
	Status      string `json:"status" gorm:"status;"`
}

type Pageing struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

func (TodoItem) TableName() string {
	return "todo_items"
}

func main() {
	dsn := "root:mypass@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	// now1 := time.Now().UTC()

	// item := TodoItem{
	// 	Id:          1,
	// 	Title:       "Title 1",
	// 	Description: "Description 1",
	// 	Status:      "Doing",
	// 	Created_at:  &now1,
	// 	Updated_at:  nil,
	// }

	// jsonData, err := json.Marshal(item)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(string(jsonData))

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("", getAll(db))
		v1.POST("", createItem(db))
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
	var result []TodoItem
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

func createItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreate
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		result := db.Table("todo_items").Create(&data)
		fmt.Println(result)

		// c.JSON(http.StatusOK, gin.H{
		// 	"id":     data.Id,
		// 	"effect": result.RowsAffected,
		// })

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))
	}
}

func getItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItem
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
		var data TodoItemUpdate
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
