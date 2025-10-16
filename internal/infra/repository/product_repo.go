package repository

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/models"
	"ecommerce/pkg/utils"
	"fmt"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(name, slug, description string, price, discountPrice float64, stock int, categoryID string, images []string) (*models.Product, error)
	GetBySlug(slug string) (*models.Product, error)
	GetById(id string) (*models.Product, error)
	List(page, limit, search string, filters map[string]interface{}) ([]*models.Product, error)
	Update(id string, updateData map[string]interface{}, images *dto.ImageUpdate) (*models.Product, error)
	SoftDelete(id string) error
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


func (r *productRepository) GetById(id string) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("Images").Find(&product).Error; err != nil {
		return nil, err
	}
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
	if err := r.db.Preload("Images").Find(&product).Error; err != nil {
		return nil, err
	}
	err := r.db.Where("slug = ? AND is_deleted = ?", slug, false).First(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) List(page, limit, search string, filters map[string]interface{}) ([]*models.Product, error) {
	var products []*models.Product

	query := r.db.Model(&models.Product{}).
		Where("is_deleted = false").
		Preload("Images")

	if search != "" {
		searchPattern := "%" + search + "%"
		query = query.Where("name ILIKE ?", searchPattern)
	}

	for k, v := range filters {
		query = query.Where(fmt.Sprintf("%s = ?", k), v)
	}

	pageInt, limitInt := utils.ParsePagination(page, limit)
	offset := (pageInt - 1) * limitInt

	err := query.Offset(offset).Limit(limitInt).Find(&products).Error
	return products, err
}



func (r *productRepository) Update(
    id string,
    updateData map[string]interface{},
    images *dto.ImageUpdate,
) (*models.Product, error) {

    // Start transaction
    tx := r.db.Begin()
    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // Update non-image fields directly
    if len(updateData) > 0 {
        if err := tx.Model(&models.Product{}).Where("id = ?", id).Updates(updateData).Error; err != nil {
            tx.Rollback()
            return nil, fmt.Errorf("failed to update product: %w", err)
        }
    }

    // Only update images if provided
    if images != nil {
        if images.Replace {
            if err := tx.Where("product_id = ?", id).Delete(&models.ProductImage{}).Error; err != nil {
                tx.Rollback()
                return nil, fmt.Errorf("failed to clear old images: %w", err)
            }
            for _, url := range images.Add {
                tx.Create(&models.ProductImage{ProductID: id, URL: url})
            }
        } else {
            if len(images.Delete) > 0 {
                if err := tx.Where("product_id = ? AND url IN ?", id, images.Delete).Delete(&models.ProductImage{}).Error; err != nil {
                    tx.Rollback()
                    return nil, fmt.Errorf("failed to delete images: %w", err)
                }
            }
            for _, url := range images.Add {
                var count int64
                tx.Model(&models.ProductImage{}).Where("product_id = ? AND url = ?", id, url).Count(&count)
                if count == 0 {
                    tx.Create(&models.ProductImage{ProductID: id, URL: url})
                }
            }
        }
    }

    // Load the updated product with images for response
    var product models.Product
    if err := tx.Preload("Images").First(&product, "id = ? AND id_deleted = ?", id).Error; err != nil {
        tx.Rollback()
        return nil, fmt.Errorf("failed to reload product: %w", err)
    }

    tx.Commit()
    return &product, nil
}

func (r *productRepository) SoftDelete(id string) error {
	return r.db.Model(&models.Product{}).Where("id = ? AND is_deleted = ?", id).Update("is_deleted", true).Error
}