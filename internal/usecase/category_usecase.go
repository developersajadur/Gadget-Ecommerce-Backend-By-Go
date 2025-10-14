package usecase

import (
	"ecommerce/internal/domain"
	"ecommerce/internal/infra/repository"
	"ecommerce/pkg/helpers"
	"fmt"
)

type CategoryUsecase interface {
	Create(name string, description string, image *string) (*domain.Category, error)
	GetBySlug(slug string) (*domain.Category, error)
	GetById(id string) (*domain.Category, error)
	List(page string, limit string, search string, filters map[string]string) ([]*domain.Category, error)
	Update(id string, name string, description string, image *string) (*domain.Category, error)
	SoftDelete(id string) error
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
		existing, err := uc.categoryRepo.GetBySlug(slug)
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

func (uc *categoryUsecase) GetBySlug(slug string) (*domain.Category, error) {
	return uc.categoryRepo.GetBySlug(slug)
}

func (uc *categoryUsecase) GetById(id string) (*domain.Category, error) {
	return uc.categoryRepo.GetById(id)
}

func (uc *categoryUsecase) List(page string, limit string, search string, filters map[string]string) ([]*domain.Category, error) {
	return uc.categoryRepo.List(page, limit, search, filters)
}

func (uc *categoryUsecase) Update(id string, name string, description string, image *string) (*domain.Category, error) {
	slug := helpers.GenerateSlug(name)
	originalSlug := slug

	suffix := 1
	for {
		existing, err := uc.categoryRepo.GetBySlug(slug)
		if err != nil && err.Error() != "sql: no rows in result set" {
			return nil, err
		}

		if existing == nil {
			break
		}
		slug = fmt.Sprintf("%s-%d", originalSlug, suffix)
		suffix++
	}
	return uc.categoryRepo.Update(id, name, slug, description, image)
}

func (uc *categoryUsecase) SoftDelete(id string) error {
	return uc.categoryRepo.SoftDelete(id)
}
