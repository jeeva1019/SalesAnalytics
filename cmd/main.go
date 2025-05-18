package main

import (
	"SalesAnalytics/config"
	"SalesAnalytics/controllers"
	"SalesAnalytics/helpers"
	"SalesAnalytics/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Setup logger
	logFile, err := helpers.SetupLogger()
	if err != nil {
		log.Fatalf("Failed to set up logger: %v", err)
	}
	defer logFile.Close()

	// Initialize TOML configuration
	config.TomlInit()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Initialize models and controllers
	salesModel := models.NewSalesModel(db)
	salesController := controllers.NewSalesController(salesModel)

	// Start automatic data refresher in background
	go AutoRefresher(salesController)

	// Setup routes
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/refresh-data", salesController.RefreshDataAPI).Methods("POST")
	api.HandleFunc("/top-products", salesController.GetTopProducts).Methods("GET")
	api.HandleFunc("/top-products/category", salesController.GetTopProductsByCategory).Methods("GET")
	api.HandleFunc("/top-products/region", salesController.GetTopProductsByRegion).Methods("GET")

	// Load port from config
	port := config.GetTomlValue("common", "port")
	if port == "" {
		port = "8080"
	}

	// Start server
	fmt.Printf("🚀 Server running on port %s...\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), r); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func AutoRefresher(salesController *controllers.SalesController) {
	debug := new(helpers.HelperStruct)
	debug.Init()

	// Run once at startup
	debug.Info("Running initial data refresh...")
	if err := salesController.RefershData(debug); err != nil {
		debug.Error("Initial data refresh failed", err)
	} else {
		debug.Info("Initial data refresh completed successfully")
	}

	hours, _ := strconv.Atoi(config.GetTomlValue("common", "hours"))
	if hours == 0 {
		hours = 10
	}
	ticker := time.NewTicker(time.Duration(hours) * time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		debug.Info("Starting scheduled 10-hour data refresh...")
		if err := salesController.RefershData(debug); err != nil {
			debug.Error("Scheduled data refresh failed", err)
		} else {
			debug.Info("Scheduled data refresh completed successfully")
		}
	}
}
