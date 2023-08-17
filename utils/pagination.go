package utils

import (
	"gorm.io/gorm"
)

var DefaultPageSize = 10
var MaxPageSize = 100

type Page struct {
	page     int
	pageSize int
}

func Paginate(page Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.page <= 0 {
			page.page = 1
		}

		switch {
		case page.pageSize > MaxPageSize:
			page.pageSize = MaxPageSize
		case page.pageSize <= 0:
			page.pageSize = DefaultPageSize
		}

		offset := (page.page - 1) * page.pageSize

		return db.Offset(offset).Limit(page.pageSize)
	}
}
