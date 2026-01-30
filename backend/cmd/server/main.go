package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"

	"github.com/paula-dot/knbs-open-data-api/backend/internal/database"
	"github.com/paula-dot/knbs-open-data-api/backend/internal/handlers"
	"github.com/paula-dot/knbs-open-data-api/backend/internal/services"
)

func main() {
	// 1. Load environment (optional .env)
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 2. Connect to DB
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	pool, err := database.NewConnection(dbURL)
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}
	defer pool.Close()

	// 3. Initialize sqlc Queries
	db := database.New(pool)

	// 4. Initialize Services
	countyService := services.NewCountyService(db)
	statsService := services.NewStatsService(db)

	// 5. Initialize Handlers
	countyHandler := handlers.NewCountyHandler(countyService)
	statsHandler := handlers.NewStatsHandler(statsService)

	// 6. Setup Router
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Essential Middleware
	r.Use(middleware.Logger) // Log every request
	r.Use(middleware.Recoverer)

	// Add top-level routes for convenience (avoid 404 when hitting /counties)
	r.Get("/counties", countyHandler.List)
	r.Get("/counties/{id}", countyHandler.GetByID)

	// 7. Define Namespaced Routes
	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/counties", countyHandler.List)
		r.Get("/counties/{id}", countyHandler.GetByID)

		// New Stats Routes
		r.Get("/indicators", statsHandler.ListIndicators)
		r.Get("/data", statsHandler.GetData)
	})

	// Build a list of registered routes for logging and 404 responses
	routesList := []string{}
	if err := chi.Walk(r, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		routesList = append(routesList, method+" "+route)
		return nil
	}); err != nil {
		log.Printf("WARN: error walking routes: %v", err)
	}

	log.Println("Registered routes:")
	for _, rt := range routesList {
		log.Println("  ", rt)
	}

	// Custom 404 handler that returns available routes to help debugging
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"error":            "not found",
			"requested_method": r.Method,
			"requested_path":   r.URL.Path,
			"available_routes": routesList,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("🚀 Server starting on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
