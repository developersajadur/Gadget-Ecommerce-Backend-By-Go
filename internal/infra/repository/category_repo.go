package repository

import (
	"ecommerce/internal/domain"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository interface {
	Create(name string, slug string, description string, image *string) (*domain.Category, error)
	FindBySlug(slug string) (*domain.Category, error)
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

func (r *categoryRepository) FindBySlug(slug string) (*domain.Category, error) {
	query := `
		SELECT id, name, slug, description, image, is_deleted, created_at, updated_at
		FROM categories
		WHERE slug = $1 AND is_deleted = FALSE
	`

	var category domain.Category
	err := r.db.Get(&category, query, slug)
	if err != nil {
		return nil, err
	}

	return &category, nil
}
