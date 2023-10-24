package business

import (
	"context"
	"strings"

	"social-todo-list/modules/item/model"
)

type CreateItemStorage interface {
	CreateItem(c context.Context, data *model.TodoItemCreate) error
}

type CreateItemBusiness struct {
	store CreateItemStorage
}

func NewCreateItem(store CreateItemStorage) *CreateItemBusiness {
	return &CreateItemBusiness{store: store}
}

func (business *CreateItemBusiness) CreateNewItem(c context.Context, data *model.TodoItemCreate) error {
	title := strings.TrimSpace(data.Title)
	if title == "" {
		return model.ErrTitleBlank
	}
	if err := business.store.CreateItem(c, data); err != nil {
		return err
	}

	return nil
}
