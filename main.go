package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type TodoItem struct {
	Id          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Created_at  *time.Time `json:"created_at"`
	Updated_at  *time.Time `json:"updated_at"`
}

func main() {
	fmt.Println("first")
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
}
