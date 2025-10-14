package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type CategoryHandler struct {
	categoryUC usecase.CategoryUsecase
}

func NewCategoryHandler(uc usecase.CategoryUsecase) *CategoryHandler {
	return &CategoryHandler{uc}
}

func (h *CategoryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Image       *string `json:"image"`
	}

	helpers.BodyDecoder(w, r, &req)
	category, err := h.categoryUC.Create(req.Name, req.Description, req.Image)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to create category")
		return
	}

	helpers.SendResponse(w, category, http.StatusOK, "Category created successfully")
}

func (h *CategoryHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	if slug == "" {
		helpers.SendError(w, nil, http.StatusBadRequest, "Slug is required")
		return
	}

	category, err := h.categoryUC.GetBySlug(slug)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch category")
		return
	}
	if category == nil {
		helpers.SendError(w, nil, http.StatusNotFound, "Category not found")
		return
	}
	helpers.SendResponse(w, category, http.StatusOK, "Category fetched successfully")
}

func (h *CategoryHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		helpers.SendError(w, nil, http.StatusBadRequest, "ID is required")
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	category, err := h.categoryUC.GetById(id)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch category")
		return
	}
	if category == nil {
		helpers.SendError(w, nil, http.StatusNotFound, "category not found")
		return
	}

	helpers.SendResponse(w, category, http.StatusOK, "Category fetched successfully")
}

func (h *CategoryHandler) List(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	page := query.Get("page")
	limit := query.Get("limit")
	search := query.Get("search")

	filters := make(map[string]string)
	for key, values := range query {
		if len(values) == 0 {
			continue
		}
		switch key {
		case "page", "limit", "search":
			continue
		default:
			filters[key] = values[0]
		}
	}

	categories, err := h.categoryUC.List(page, limit, search, filters)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch categories")
		return
	}

	helpers.SendResponse(w, categories, http.StatusOK, "Categories fetched successfully")
}

func (h *CategoryHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		helpers.SendError(w, nil, http.StatusBadRequest, "ID is required")
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Image       *string `json:"image"`
	}

	helpers.BodyDecoder(w, r, &req)

	category, err := h.categoryUC.Update(id, req.Name, req.Description, req.Image)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to update category")
		return
	}
	if category == nil {
		helpers.SendError(w, nil, http.StatusNotFound, "Category not found")
		return
	}
	helpers.SendResponse(w, category, http.StatusOK, "Category updated successfully")
}

// SoftDelete handles soft deleting a category
func (h *CategoryHandler) SoftDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		helpers.SendError(w, nil, http.StatusBadRequest, "ID is required")
		return
	}

	if _, err := uuid.Parse(id); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Invalid UUID format")
		return
	}

	err := h.categoryUC.SoftDelete(id)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to delete category")
		return
	}

	helpers.SendResponse(w, nil, http.StatusOK, "Category deleted successfully")
}
