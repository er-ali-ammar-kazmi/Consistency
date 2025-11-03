package util

import (
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

func Paginate(req *http.Request) func(tx *gorm.DB) *gorm.DB {

	page, _ := strconv.Atoi(req.URL.Query().Get("page"))
	if page < 1 {
		page = 1
	} else if page > 5 {
		page = 5
	}

	pageSize := 5
	return func(tx *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return tx.Offset(offset).Limit(pageSize)
	}
}
