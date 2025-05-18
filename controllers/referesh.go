package controllers

import (
	"SalesAnalytics/config"
	"SalesAnalytics/constants"
	"SalesAnalytics/helpers"
	"SalesAnalytics/models"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/gorm/clause"
)

type SalesController struct {
	models *models.SalesModel
}

func NewSalesController(models *models.SalesModel) *SalesController {
	return &SalesController{models}
}

func (s *SalesController) RefreshDataAPI(w http.ResponseWriter, r *http.Request) {
	debug := new(helpers.HelperStruct)
	debug.Init()
	debug.Info("RefreshDataAPI(+)")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, credentials")

	if strings.EqualFold(r.Method, "POST") {
		if err := s.RefershData(debug); err != nil {
			debug.Error("RDA:001", err)
			fmt.Fprintln(w, ResponseConstructor(constants.ERROR, "RDA:001 "+err.Error(), nil))
			return
		}
		fmt.Fprintln(w, ResponseConstructor(constants.SUCCESS, "Process successfully completed", nil))
	}
}

func (s *SalesController) RefershData(debug *helpers.HelperStruct) error {
	debug.Info("RefershData(+)")

	csvPath := config.GetTomlValue("common", "path")
	file, err := os.Open(csvPath)
	if err != nil {
		debug.Error("CRRD:001", err.Error())
		return err
	}
	defer file.Close()

	var salesRcds []*models.SalesData
	err = CSVFileReader(debug, file, &salesRcds)
	if err != nil {
		debug.Error("CRRD:002", err)
		return err
	}

	var customers []models.Customer
	var products []models.Product
	var orders []models.Order
	var orderDetails []models.OrderDetail

	for _, salesRcd := range salesRcds {
		var sp models.SalesProcessData

		sp.Customer.ID = salesRcd.CustomerID
		sp.Customer.Name = salesRcd.CustomerName
		sp.Customer.Email = salesRcd.CustomerEmail
		sp.Customer.Address = salesRcd.CustomerAddress
		customers = append(customers, sp.Customer)

		sp.Product.ProductID = salesRcd.ProductID
		sp.Product.ProductName = salesRcd.ProductName
		sp.Product.Category = salesRcd.Category
		products = append(products, sp.Product)

		sp.Order.OrderID = salesRcd.OrderID
		sp.Order.CustomerID = salesRcd.CustomerID
		sp.Order.DateOfSale, _ = time.Parse("2006-01-02", salesRcd.DateOfSale)
		sp.Order.PaymentMethod = salesRcd.PaymentMethod
		sp.Order.Region = salesRcd.Region
		sp.Order.ShippingCost, _ = strconv.ParseFloat(salesRcd.ShippingCost, 64)
		orders = append(orders, sp.Order)

		sp.OrderDetail.OrderID = salesRcd.OrderID
		sp.OrderDetail.ProductID = salesRcd.ProductID
		sp.OrderDetail.QuantitySold, _ = strconv.Atoi(salesRcd.QuantitySold)
		sp.OrderDetail.UnitPrice = salesRcd.UnitPrice
		sp.OrderDetail.Discount = salesRcd.Discount
		orderDetails = append(orderDetails, sp.OrderDetail)
	}

	s.models.DB.Clauses(OnConflictUpdateAll()).Create(&customers)
	s.models.DB.Clauses(OnConflictUpdateAll()).Create(&products)
	s.models.DB.Clauses(OnConflictUpdateAll()).Create(&orders)
	s.models.DB.Clauses(OnConflictUpdateAll()).Create(&orderDetails)

	debug.Info("RefershData(-)")
	return nil
}

// OnConflictUpdateAll returns a clause to update on conflict (upsert)
func OnConflictUpdateAll() clause.OnConflict {
	return clause.OnConflict{
		UpdateAll: true,
	}
}
