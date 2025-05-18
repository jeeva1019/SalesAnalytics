package models

import "time"

type Customer struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Email   string
	Address string

	Orders []Order `gorm:"foreignKey:CustomerID"`
}

type Product struct {
	ProductID   string `gorm:"primaryKey"`
	ProductName string
	Category    string

	OrderDetails []OrderDetail `gorm:"foreignKey:ProductID"`
}

type Order struct {
	OrderID       uint   `gorm:"primaryKey"`
	CustomerID    string `gorm:"index"`
	DateOfSale    time.Time
	Region        string
	ShippingCost  float64
	PaymentMethod string

	Customer     Customer      `gorm:"foreignKey:CustomerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	OrderDetails []OrderDetail `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type OrderDetail struct {
	OrderID      uint   `gorm:"primaryKey"`
	ProductID    string `gorm:"primaryKey"`
	QuantitySold int
	UnitPrice    float64
	Discount     float64

	Product Product `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
	Order   Order   `gorm:"foreignKey:OrderID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

type SalesData struct {
	OrderID         uint    `csv:"Order ID"`
	ProductID       string  `csv:"Product ID"`
	CustomerID      string  `csv:"Customer ID"`
	ProductName     string  `csv:"Product Name"`
	Category        string  `csv:"Category"`
	Region          string  `csv:"Region"`
	DateOfSale      string  `csv:"Date of Sale"`
	QuantitySold    string  `csv:"Quantity Sold"`
	UnitPrice       float64 `csv:"Unit Price"`
	Discount        float64 `csv:"Discount"`
	ShippingCost    string  `csv:"Shipping Cost"`
	PaymentMethod   string  `csv:"Payment Method"`
	CustomerName    string  `csv:"Customer Name"`
	CustomerEmail   string  `csv:"Customer Email"`
	CustomerAddress string  `csv:"Customer Address"`
}

type SalesProcessData struct {
	Customer
	Product
	Order
	OrderDetail
}
