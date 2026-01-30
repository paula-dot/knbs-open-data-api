package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/paula-dot/knbs-open-data-api/backend/internal/services"
)

type CountyHandler struct {
	service services.CountyService
}

func NewCountyHandler(service services.CountyService) *CountyHandler {
	return &CountyHandler{
		service: service,
	}
}

// List handles GET /counties
func (h *CountyHandler) List(w http.ResponseWriter, r *http.Request) {
	counties, err := h.service.GetAllCounties(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch counties", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": counties,
	})
}

func (h *CountyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// 1. Extract ID from URL
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid county ID", http.StatusBadRequest)
		return
	}

	// 2. Call service to get county by ID
	county, err := h.service.GetCountyByID(r.Context(), int32(id))
	if err != nil {
		http.Error(w, "County not found", http.StatusNotFound)
		return
	}

	// 3. Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": county,
	})
}
