package model

type CSVRow struct {
	OrderID       string  `csv:"Order ID"`
	ProductID     string  `csv:"Product ID"`
	CustomerID    string  `csv:"Customer ID"`
	ProductName   string  `csv:"Product Name"`
	Category      string  `csv:"Category"`
	Region        string  `csv:"Region"`
	DateOfSale    string  `csv:"Date of Sale"`
	QuantitySold  int     `csv:"Quantity Sold"`
	UnitPrice     float64 `csv:"Unit Price"`
	Discount      float64 `csv:"Discount"`
	ShippingCost  float64 `csv:"Shipping Cost"`
	PaymentMethod string  `csv:"Payment Method"`
	CustomerName  string  `csv:"Customer Name"`
	CustomerEmail string  `csv:"Customer Email"`
	CustomerAddr  string  `csv:"Customer Address"`
}
