package querybuilder

import (
	"gorm.io/gorm"
	"strings"
)

type GormQueryBuilder struct {
	DB       *gorm.DB
	Filters  map[string]interface{}
	Search   string
	SearchCols []string
	Page     int
	Limit    int
	OrderBy  string
}

func New(db *gorm.DB) *GormQueryBuilder {
	return &GormQueryBuilder{
		DB:       db,
		Filters:  map[string]interface{}{},
		Page:     1,
		Limit:    20,
		OrderBy:  "created_at desc",
	}
}

func (qb *GormQueryBuilder) SetFilters(filters map[string]interface{}) {
	for k, v := range filters {
		qb.Filters[k] = v
	}
}

func (qb *GormQueryBuilder) SetSearch(search string, cols []string) {
	qb.Search = search
	qb.SearchCols = cols
}

func (qb *GormQueryBuilder) SetPagination(page, limit int) {
	if page > 0 {
		qb.Page = page
	}
	if limit > 0 {
		qb.Limit = limit
	}
}

func (qb *GormQueryBuilder) SetOrder(order string) {
	if order != "" {
		qb.OrderBy = order
	}
}

func (qb *GormQueryBuilder) Build() *gorm.DB {
	db := qb.DB

	// Filters
	for k, v := range qb.Filters {
		db = db.Where(k+" = ?", v)
	}

	// Search
	if qb.Search != "" && len(qb.SearchCols) > 0 {
		var searchParts []string
		var args []interface{}
		for _, col := range qb.SearchCols {
			searchParts = append(searchParts, col+" ILIKE ?")
			args = append(args, "%"+qb.Search+"%")
		}
		db = db.Where(strings.Join(searchParts, " OR "), args...)
	}

	// Pagination
	offset := (qb.Page - 1) * qb.Limit
	db = db.Offset(offset).Limit(qb.Limit)

	// Order
	db = db.Order(qb.OrderBy)

	return db
}
