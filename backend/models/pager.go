package models

import "math"

type Pager struct {
	Count       int64       `json:"count"`
	CurrentPage int64       `json:"current_page"`
	Data        interface{} `json:"data"`
	PageSize    int64       `json:"page_size"`
	TotalPages  int64       `json:"total_pages"`
}

func PageCount(count int64, pageSize int64, page int64) (totalPage int64, offset int64, currentPage int64) {
	pageCount := int64(math.Ceil(float64(count) / float64(pageSize)))
	if page > pageCount {
		page = pageCount
	}
	// Handle edge case when page is 0 or less
	if page < 1 {
		return pageCount, 0, page
	}
	return pageCount, pageSize * (page - 1), page
}
