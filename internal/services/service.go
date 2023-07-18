package services

import (
	"encoding/json"
	"net/http"
	"portal_service/internal/models"
)

// TODO add application logic
func getProducts(w http.ResponseWriter, r *http.Request) {
	var User []models.User
	result := h.db.Preload("Category").Find(&User)
	if result.Error != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	enc.SetIndent("", "    ")
	enc.Encode(User)
}
