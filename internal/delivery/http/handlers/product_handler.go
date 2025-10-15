package handlers

import (
	"ecommerce/internal/dto"
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	productUC usecase.ProductUsecase
}

func NewProductHandler(uc usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc}
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name          string   `json:"name"`
		Description   string   `json:"description"`
		Price         float64  `json:"price"`
		DiscountPrice float64  `json:"discount_price"`
		Stock         int      `json:"stock"`
		CategoryID    string   `json:"category_id"`
		Images        []string `json:"images"`
	}

	// Decode the request body
	if err := helpers.BodyDecoder(w, r, &req); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Invalid request body")
		return
	}

	// Create the product
	product, err := h.productUC.Create(
		req.Name,
		req.Description,
		req.Price,
		req.DiscountPrice,
		req.Stock,
		req.CategoryID,
		req.Images,
	)
	fmt.Println(err)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to create product")
		return
	}

	// Send the response
	helpers.SendResponse(w, product, http.StatusOK, "Product created successfully")
}

func (h *ProductHandler) GetBySlug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]
	if slug == "" {
		helpers.SendError(w, fmt.Errorf("slug is required"), http.StatusBadRequest, "Slug parameter is missing")
		return
	}

	product, err := h.productUC.GetBySlug(slug)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch product")
		return
	}
	if product == nil {
		helpers.SendError(w, fmt.Errorf("product not found"), http.StatusNotFound, "Product not found")
		return
	}

	helpers.SendResponse(w, product, http.StatusOK, "Product fetched successfully")
}

func (h *ProductHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		helpers.SendError(w, fmt.Errorf("id is required"), http.StatusBadRequest, "ID parameter is missing")
		return
	}

	product, err := h.productUC.GetById(id)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to fetch product")
		return
	}
	if product == nil {
		helpers.SendError(w, fmt.Errorf("product not found"), http.StatusNotFound, "Product not found")
		return
	}

	helpers.SendResponse(w, product, http.StatusOK, "Product fetched successfully")
}
func (h *ProductHandler) List(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	search := r.URL.Query().Get("search")
	filters := make(map[string]string)
	for key, values := range r.URL.Query() {
		if len(values) > 0 && key != "page" && key != "limit" && key != "search" {
			filters[key] = values[0]
		}
	}

	products, err := h.productUC.List(page, limit, search, filters)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to list products")
		return
	}

	helpers.SendResponse(w, products, http.StatusOK, "Products listed successfully")
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		helpers.SendError(w, fmt.Errorf("id is required"), http.StatusBadRequest, "ID parameter is missing")
		return
	}

	var req struct {
		UpdateData map[string]interface{} `json:"update_data"`
		Images     *dto.ImageUpdate      `json:"images"`
	}

	if err := helpers.BodyDecoder(w, r, &req); err != nil {
		helpers.SendError(w, err, http.StatusBadRequest, "Invalid request body")
		return
	}

	product, err := h.productUC.Update(id, req.UpdateData, req.Images)
	if err != nil {
		helpers.SendError(w, err, http.StatusInternalServerError, "Failed to update product")
		return
	}

	helpers.SendResponse(w, product, http.StatusOK, "Product updated successfully")
}


func (h *ProductHandler) SoftDelete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    if id == "" {
        helpers.SendError(w, fmt.Errorf("id is required"), http.StatusBadRequest, "ID parameter is missing")
        return
    }

    if err := h.productUC.SoftDelete(id); err != nil {
        helpers.SendError(w, err, http.StatusInternalServerError, "Failed to delete product")
        return
    }

    helpers.SendResponse(w, nil, http.StatusOK, "Product deleted successfully")
}