package usecase

import (
	"errors"
	"fmt"

	"ecommerce/internal/infra/repository"
	"ecommerce/internal/models"
	"ecommerce/pkg/helpers"
)

type ProductUsecase interface {
	Create(name, description string, price, discountPrice float64, stock int, categoryID string, images []string) (*models.Product, error)
	GetBySlug(slug string) (*models.Product, error)
	GetById(id string) (*models.Product, error)
	List(page string, limit string, search string, filters map[string]string) ([]*models.Product, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
	categoryUC  CategoryUsecase
}

func NewProductUsecase(productRepo repository.ProductRepository, categoryUC CategoryUsecase) ProductUsecase {
	return &productUsecase{
		productRepo: productRepo,
		categoryUC:  categoryUC,
	}
}

func (uc *productUsecase) Create(
	name, description string,
	price, discountPrice float64,
	stock int,
	categoryID string,
	images []string,
) (*models.Product, error) {
	// Validate category existence
	category, err := uc.categoryUC.GetById(categoryID)
	if err != nil {
		return nil, fmt.Errorf("failed to check category: %w", err)
	}
	if category == nil {
		return nil, errors.New("category not found")
	}

	// Generate unique slug
	slug := helpers.GenerateSlug(name)
	originalSlug := slug
	suffix := 1

	for {
		existing, err := uc.productRepo.GetBySlug(slug)
		if err != nil {
			return nil, err
		}
		if existing == nil {
			break
		}
		slug = fmt.Sprintf("%s-%d", originalSlug, suffix)
		suffix++
	}

	// Create product via repository
	return uc.productRepo.Create(name, slug, description, price, discountPrice, stock, categoryID, images)
}

func (uc *productUsecase) GetBySlug(slug string) (*models.Product, error) {
	product, err := uc.productRepo.GetBySlug(slug)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}

func (uc *productUsecase) GetById(id string) (*models.Product, error) {
	product, err := uc.productRepo.GetById(id)
	if err != nil {
		return nil, err
	}
	if product == nil {
		return nil, errors.New("product not found")
	}
	return product, nil
}


func (uc *productUsecase) List(page string, limit string, search string, filters map[string]string) ([]*models.Product, error) {
	convertedFilters := make(map[string]interface{}, len(filters))
	for k, v := range filters {
		convertedFilters[k] = v
	}
	return uc.productRepo.List(page, limit, search, convertedFilters)
}