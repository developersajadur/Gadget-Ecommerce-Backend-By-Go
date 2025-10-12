package usecase

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/helpers"
	"fmt"
)

type CategoryUsecase interface {
	Create(name string, description string, image *string) (*domain.Category, error)
}

type categoryUsecase struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo}
}

func (uc *categoryUsecase) Create(name, description string, image *string) (*domain.Category, error) {
	slug := helpers.GenerateSlug(name)
	originalSlug := slug

	suffix := 1
	for {
		existing, err := uc.categoryRepo.FindBySlug(slug)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, err
		}

		if existing == nil {
			break
		}
		slug = fmt.Sprintf("%s-%d", originalSlug, suffix)
		suffix++
	}

	return uc.categoryRepo.Create(name, slug, description, image)
}
