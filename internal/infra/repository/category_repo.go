package repository

import (
	"database/sql"
	"ecommerce/internal/domain"
	"ecommerce/pkg/querybuilder"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	Create(name string, slug string, description string, image *string) (*domain.Category, error)
	GetBySlug(slug string) (*domain.Category, error)
	GetById(id string) (*domain.Category, error)
	List(page string, limit string, search string, filters map[string]string) ([]*domain.Category, error)
	Update(id string, name string, slug string, description string, image *string) (*domain.Category, error)
	SoftDelete(id string) error
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(name, slug, description string, image *string) (*domain.Category, error) {
	query := `
		INSERT INTO categories (name, slug, description, image)
		VALUES ($1, $2, $3, $4)
		RETURNING id, name, slug, description, image, is_deleted, created_at, updated_at
	`

	var category domain.Category
	err := r.db.QueryRowx(query, name, slug, description, image).StructScan(&category)
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) GetBySlug(slug string) (*domain.Category, error) {
	var category domain.Category
	query := `SELECT * FROM categories WHERE slug = $1 AND is_deleted = FALSE`
	err := r.db.Get(&category, query, slug)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetById(id string) (*domain.Category, error) {
	var category domain.Category
	query := `SELECT * FROM categories WHERE id = $1 AND is_deleted = FALSE`
	err := r.db.Get(&category, query, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) List(page, limit, search string, filters map[string]string) ([]*domain.Category, error) {
	pageInt, _ := strconv.Atoi(page)
	limitInt, _ := strconv.Atoi(limit)

	qb := querybuilder.New("SELECT * FROM categories")
	qb.SetPagination(pageInt, limitInt)
	qb.SetSearch(search, []string{"name"})
	qb.AddFilters(filters)

	query, args := qb.Build()

	var categories []domain.Category
	if err := r.db.Select(&categories, query, args...); err != nil {
		return nil, err
	}

	categoryPtrs := make([]*domain.Category, len(categories))
	for i := range categories {
		categoryPtrs[i] = &categories[i]
	}

	return categoryPtrs, nil
}

func (r *categoryRepository) Update(id, name, slug, description string, image *string) (*domain.Category, error) {
	query := `
		UPDATE categories
		SET name = $1, slug = $2, description = $3, image = $4, updated_at = NOW()
		WHERE id = $5 AND is_deleted = FALSE
		RETURNING id, name, slug, description, image, is_deleted, created_at, updated_at
	`

	var category domain.Category
	err := r.db.QueryRowx(query, name, slug, description, image, id).StructScan(&category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &category, nil
}

func (r *categoryRepository) SoftDelete(id string) error {
	query := `UPDATE categories SET is_deleted = TRUE, updated_at = NOW() WHERE id = $1 AND is_deleted = FALSE`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
