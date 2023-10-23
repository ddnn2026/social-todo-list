package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
}
type TodoItemCreate struct {
	Id          int    `json:"-" gorm:"id;"`
	Title       string `json:"title" gorm:"title;"`
	Description string `json:"description" gorm:"description;"`
	Status      string `json:"status" gorm:"status;"`
}

func main() {
	dsn := "root:mypass@tcp(127.0.0.1:3306)/mydb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(db)

	now1 := time.Now().UTC()

	item := TodoItem{
		Id:          1,
		Title:       "Title 1",
		Description: "Description 1",
		Status:      "Doing",
		Created_at:  &now1,
		Updated_at:  nil,
	}

	jsonData, err := json.Marshal(item)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))

	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("")
		v1.POST("", CreateItem(db))
		v1.GET("/:id")
		v1.PATCH("/:id")
		v1.DELETE("/:id")
	}

	r.GET("/ping", getPing)
	r.Run(":8088")
}

func getPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateItem(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		var data TodoItemCreate
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		result := db.Table("todo_items").Create(&data)
		fmt.Println(result)

		c.JSON(http.StatusOK, gin.H{
			"id":     data.Id,
			"effect": result.RowsAffected,
		})
	}
}
