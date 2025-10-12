package querybuilder

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	BaseQuery string
	Filters   map[string]string
	Search    string
	SearchCols []string
	Limit     int
	Offset    int
	OrderBy   string
	Args      []interface{}
	argCount  int
}

func New(baseQuery string) *QueryBuilder {
	return &QueryBuilder{
		BaseQuery: baseQuery,
		Filters:   make(map[string]string),
		Args:      []interface{}{},
		argCount:  1,
	}
}

func (qb *QueryBuilder) SetPagination(page, limit int) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 20
	}
	qb.Limit = limit
	qb.Offset = (page - 1) * limit
}

func (qb *QueryBuilder) SetSearch(search string, columns []string) {
	qb.Search = search
	qb.SearchCols = columns
}

func (qb *QueryBuilder) AddFilters(filters map[string]string) {
	for k, v := range filters {
		qb.Filters[k] = v
	}
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	var whereParts []string

	// Handle filters
	for key, val := range qb.Filters {
		whereParts = append(whereParts, fmt.Sprintf("%s = $%d", key, qb.argCount))
		qb.Args = append(qb.Args, val)
		qb.argCount++
	}

	// Handle search
	if qb.Search != "" && len(qb.SearchCols) > 0 {
		var searchParts []string
		for _, col := range qb.SearchCols {
			searchParts = append(searchParts, fmt.Sprintf("%s ILIKE $%d", col, qb.argCount))
		}
		qb.Args = append(qb.Args, "%"+qb.Search+"%")
		whereParts = append(whereParts, "("+strings.Join(searchParts, " OR ")+")")
		qb.argCount++
	}

	query := qb.BaseQuery
	if len(whereParts) > 0 {
		query += " WHERE " + strings.Join(whereParts, " AND ")
	}

	if qb.OrderBy != "" {
		query += fmt.Sprintf(" ORDER BY %s", qb.OrderBy)
	} else {
		query += " ORDER BY created_at DESC"
	}

	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", qb.argCount, qb.argCount+1)
	qb.Args = append(qb.Args, qb.Limit, qb.Offset)

	return query, qb.Args
}
