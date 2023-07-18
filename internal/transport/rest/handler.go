package rest

import (
	"net/http"
	"portal_service/internal/services"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// TODO add application handler
func NewHandler(db *gorm.DB) http.Handler {

	r := mux.NewRouter().StrictSlash(true)
	// r.HandleFunc("/seed", h.seedDatabase).Methods(http.MethodGet)
	r.HandleFunc("/products", services).Methods(http.MethodGet)
	// r.HandleFunc("/products/{id}", h.getProduct).Methods(http.MethodGet)
	// r.HandleFunc("/categories", h.getCategories).Methods(http.MethodGet)
	// r.HandleFunc("/categories/{id}", h.getCategory).Methods(http.MethodGet)

	return r
}
