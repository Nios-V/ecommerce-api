package products

import (
	"net/http"

	"github.com/Nios-V/ecommerce/api/internal/products/json"
)

type handler struct {
	service Service
}

func NewHandler(svc Service) *handler {
	return &handler{
		service: svc,
	}
}

func (h *handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	err := h.service.GetAllProducts(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	products := []string{"Product 1", "Product 2", "Product 3"}
	json.Write(w, http.StatusOK, products)
}
