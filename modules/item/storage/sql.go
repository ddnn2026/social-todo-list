package storage

import "gorm.io/gorm"

type SQLstorage struct {
	db *gorm.DB
}

func NewSQLStore(db *gorm.DB) *SQLstorage {
	return &SQLstorage{
		db: db,
	}
}
