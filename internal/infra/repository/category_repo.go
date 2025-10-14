package repository

import (
	"ecommerce/internal/models"
	"ecommerce/pkg/utils"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(name, slug string, description string, image *string) (*models.Category, error)
	GetBySlug(slug string) (*models.Category, error)
	GetById(id string) (*models.Category, error)
	List(page, limit, search string, filters map[string]interface{}) ([]*models.Category, error)
	Update(id, name, slug, description string, image *string) (*models.Category, error)
	SoftDelete(id string) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) Create(name, slug string, description string, image *string) (*models.Category, error) {
	category := &models.Category{
		Name:        name,
		Slug:        slug,
		Description: &description,
		Image:       image,
	}

	if err := r.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (r *categoryRepository) GetBySlug(slug string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("slug = ? AND is_deleted = false", slug).First(&category).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &category, err
}

func (r *categoryRepository) GetById(id string) (*models.Category, error) {
	var category models.Category
	err := r.db.Where("id = ? AND is_deleted = false", id).First(&category).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &category, err
}

func (r *categoryRepository) List(page, limit, search string, filters map[string]interface{}) ([]*models.Category, error) {
	var categories []*models.Category
	query := r.db.Model(&models.Category{}).Where("is_deleted = false")

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("name ILIKE ?", searchPattern)
	}

	for k, v := range filters {
		query = query.Where(k+" = ?", v)
	}

	pageInt, limitInt := utils.ParsePagination(page, limit)
	offset := (pageInt - 1) * limitInt

	err := query.Offset(offset).Limit(limitInt).Find(&categories).Error
	return categories, err
}

func (r *categoryRepository) Update(id, name, slug, description string, image *string) (*models.Category, error) {
	category, err := r.GetById(id)
	if err != nil {
		return nil, err
	}
	if category == nil {
		return nil, nil
	}

	category.Name = name
	category.Slug = slug
	category.Description = &description
	category.Image = image

	if err := r.db.Save(category).Error; err != nil {
		return nil, err
	}

	return category, nil
}

func (r *categoryRepository) SoftDelete(id string) error {
	result := r.db.Model(&models.Category{}).Where("id = ? AND is_deleted = false", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"updated_at": time.Now(),
		})
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return result.Error
}
