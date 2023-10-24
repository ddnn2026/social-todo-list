package model

import (
	"errors"

	"social-todo-list/common"
)

var ErrTitleBlank = errors.New("title cannot be blank")

func (TodoItem) TableName() string {
	return "todo_items"
}

func (TodoItemCreate) TableName() string {
	return TodoItem{}.TableName()
}

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
