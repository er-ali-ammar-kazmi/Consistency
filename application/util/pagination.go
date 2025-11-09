package util

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(req *http.Request) func(tx *gorm.DB) *gorm.DB {

	queryString := req.URL.Query()
	page, _ := strconv.Atoi(queryString.Get("page"))
	if page < 1 {
		page = 1
	} else if page > 5 {
		page = 5
	}

	pageSize, _ := strconv.Atoi(queryString.Get("page_size"))
	if pageSize < 5 {
		pageSize = 5
	} else if page > 25 {
		pageSize = 25
	}

	return func(tx *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return tx.Offset(offset).Limit(pageSize)
	}
}
