package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"net/http"
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