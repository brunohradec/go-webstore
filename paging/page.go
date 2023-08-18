package paging

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DefaultPageSize = 10
var MaxPageSize = 100

type Page struct {
	Page     int
	PageSize int
}

func Paginate(page Page) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page.Page <= 0 {
			page.Page = 1
		}

		switch {
		case page.PageSize > MaxPageSize:
			page.PageSize = MaxPageSize
		case page.PageSize <= 0:
			page.PageSize = DefaultPageSize
		}

		offset := (page.Page - 1) * page.PageSize

		return db.Offset(offset).Limit(page.PageSize)
	}
}

func ParsePageFromQuery(c *gin.Context) Page {
	page, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		log.Printf("Page index could not be parsed from query string. Defaulting to 0")
		page = 0
	}

	pageSize, err := strconv.Atoi(c.Query("page"))

	if err != nil {
		log.Printf(
			"Page size could not be parsed from query string. Defaulting to %d",
			DefaultPageSize,
		)

		page = DefaultPageSize
	}

	return Page{
		Page:     page,
		PageSize: pageSize,
	}
}
