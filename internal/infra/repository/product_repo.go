package repository

import (
	"ecommerce/internal/models"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(name, slug, description string, price, discountPrice float64, stock int, categoryID string, images []string) (*models.Product, error)
	GetBySlug(slug string) (*models.Product, error)
	GetById(id string) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Create inserts a new product
func (r *productRepository) Create(
	name, slug, description string,
	price, discountPrice float64,
	stock int,
	categoryID string,
	images []string,
) (*models.Product, error) {
	// Convert []string to []models.ProductImage
	var productImages []models.ProductImage
	for _, img := range images {
		productImages = append(productImages, models.ProductImage{URL: img})
	}

	product := &models.Product{
		Name:          name,
		Slug:          slug,
		Description:   description,
		Price:         price,
		DiscountPrice: &discountPrice,
		Stock:         stock,
		CategoryID:    &categoryID,
		Images:        productImages,
		IsPublished:   true,
	}

	err := r.db.Create(product).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create product: %w", err)
	}

	return product, nil
}

// GetById fetches a product by ID, excluding soft-deleted ones
func (r *productRepository) GetById(id string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("id = ? AND is_deleted = ?", id, false).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

// GetBySlug fetches a product by slug, excluding soft-deleted ones
func (r *productRepository) GetBySlug(slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.Where("slug = ? AND is_deleted = ?", slug, false).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}
