package utils

import (
	"gorm.io/gorm"
)

var DefaultPageSize = 10
var MaxPageSize = 100

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > MaxPageSize:
			pageSize = MaxPageSize
		case pageSize <= 0:
			pageSize = DefaultPageSize
		}

		offset := (page - 1) * pageSize

		return db.Offset(offset).Limit(pageSize)
	}
}
