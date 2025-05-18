package controllers

import (
	"SalesAnalytics/constants"
	"SalesAnalytics/helpers"
	"fmt"
	"net/http"
	"strconv"
)

func (s *SalesController) GetTopProducts(w http.ResponseWriter, r *http.Request) {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.Info("GetTopProducts(+)")

	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	if err := Validator(map[string]string{
		"start": start,
		"end":   end,
	}); err != nil {
		debug.Error("CGTP:001", "Validation failed", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "CGTP:001 "+err.Error(), nil))
		return
	}

	var results []struct {
		ProductID   string
		ProductName string
		Quantity    int
	}

	if err := s.models.DB.
		Table("order_details").
		Select("products.product_id, products.product_name, SUM(order_details.quantity_sold) as quantity").
		Joins("JOIN products ON products.product_id = order_details.product_id").
		Joins("JOIN orders ON orders.order_id = order_details.order_id").
		Where("orders.date_of_sale BETWEEN ? AND ?", start, end).
		Group("products.product_id").
		Order("quantity DESC").
		Limit(limit).
		Scan(&results).Error; err != nil {
		debug.Error("CGTP:002", "Failed to fetch data", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "CGTP:002 "+err.Error(), nil))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, ResponseConstructor(constants.SUCCESS, http.StatusText(http.StatusOK), results))
	debug.Info("GetTopProducts(-)")
}

func (s *SalesController) GetTopProductsByCategory(w http.ResponseWriter, r *http.Request) {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.Info("GetTopProductsByCategory(+)")

	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	category := r.URL.Query().Get("category")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	if err := Validator(map[string]string{
		"start":    start,
		"end":      end,
		"category": category,
	}); err != nil {
		debug.Error("CGTPBC:001", "Validation failed", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "CGTPBC:001 "+err.Error(), nil))
		return
	}

	var results []struct {
		ProductID   string
		ProductName string
		Quantity    int
	}

	if err := s.models.DB.
		Table("order_details").
		Select("products.product_id, products.product_name, SUM(order_details.quantity_sold) as quantity").
		Joins("JOIN products ON products.product_id = order_details.product_id").
		Joins("JOIN orders ON orders.order_id = order_details.order_id").
		Where("orders.date_of_sale BETWEEN ? AND ? AND products.category = ?", start, end, category).
		Group("products.product_id").
		Order("quantity DESC").
		Limit(limit).
		Scan(&results).Error; err != nil {
		debug.Error("CGTPBC:002", "Failed to fetch data", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "CGTPBC:002 "+err.Error(), results))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, ResponseConstructor(constants.SUCCESS, http.StatusText(http.StatusOK), results))
	debug.Info("GetTopProductsByCategory(+)")
}

func (s *SalesController) GetTopProductsByRegion(w http.ResponseWriter, r *http.Request) {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.Info("GetTopProductsByRegion(+)")

	// CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	start := r.URL.Query().Get("start")
	end := r.URL.Query().Get("end")
	region := r.URL.Query().Get("region")
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	if limit == 0 {
		limit = 10
	}

	if err := Validator(map[string]string{
		"start":  start,
		"end":    end,
		"region": region,
	}); err != nil {
		debug.Error("CGTPBR:001", "Validation failed", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "CGTPBR:001 "+err.Error(), nil))
		return
	}

	var results []struct {
		ProductID   string
		ProductName string
		Quantity    int
	}

	if err := s.models.DB.
		Table("order_details").
		Select("products.product_id, products.product_name, SUM(order_details.quantity_sold) as quantity").
		Joins("JOIN products ON products.product_id = order_details.product_id").
		Joins("JOIN orders ON orders.order_id = order_details.order_id").
		Where("orders.date_of_sale BETWEEN ? AND ? AND orders.region = ?", start, end, region).
		Group("products.product_id").
		Order("quantity DESC").
		Limit(limit).
		Scan(&results).Error; err != nil {
		debug.Error("CGTPBR:002", "Failed to fetch data", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "CGTPBR:002 "+err.Error(), nil))
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, ResponseConstructor(constants.SUCCESS, http.StatusText(http.StatusOK), results))
	debug.Info("GetTopProductsByRegion(-)")
}
