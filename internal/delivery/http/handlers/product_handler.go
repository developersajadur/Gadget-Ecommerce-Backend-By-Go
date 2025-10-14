package handlers

import (
	"ecommerce/internal/usecase"
	"ecommerce/pkg/helpers"
	"fmt"
	"net/http"
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
