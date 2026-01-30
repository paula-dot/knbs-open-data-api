package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/paula-dot/knbs-open-data-api/backend/internal/services"
)

type StatsHandler struct {
	service services.StatsService
}

func NewStatsHandler(service services.StatsService) *StatsHandler {
	return &StatsHandler{
		service: service,
	}
}

// ListIndicators handles GET /indicators
func (h *StatsHandler) GetIndicators(w http.ResponseWriter, r *http.Request) {
	indicators, err := h.service.GetIndicators(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch indicators", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": indicators,
	})
}

// GetData handles GET /data?indicator=POP_TOTAL&year=2019
func (h *StatsHandler) GetData(w http.ResponseWriter, r *http.Request) {
	indicatorCode := r.URL.Query().Get("indicator")
	yearStr := r.URL.Query().Get("year")

	if indicatorCode == "" || yearStr == "" {
		http.Error(w, "Missing 'indicator' or 'year' query params", http.StatusBadRequest)
		return
	}

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	data, err := h.service.GetData(r.Context(), indicatorCode, int32(year))
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": data,
		"meta": map[string]interface{}{
			"indicator": indicatorCode,
			"year":      year,
		},
	})

}

func (h *StatsHandler) ListIndicators(writer http.ResponseWriter, request *http.Request) {

}
