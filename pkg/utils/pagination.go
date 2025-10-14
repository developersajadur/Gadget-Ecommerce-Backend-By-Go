package utils

import "strconv"


func ParsePagination(pageStr, limitStr string) (page, limit int) {
	page = 1  
	limit = 20 

	if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
		page = p
	}

	if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
		limit = l
	}

	return page, limit
}
