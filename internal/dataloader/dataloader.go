package dataloader

import (
	"lumel/internal/model"
	"lumel/internal/repo"
	"os"
	"time"

	"github.com/gocarina/gocsv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var batchSize = 100

func LoadSalesData() {
	file, err := os.Open("sample.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rows []*model.CSVRow
	if err := gocsv.UnmarshalFile(file, &rows); err != nil {
		panic(err)
	}

	customerMap := make(map[string]primitive.ObjectID)
	productMap := make(map[string]primitive.ObjectID)
	var customers []model.Customer
	var products []model.Product
	var orders []model.Order

	// Process in batches of 100
	for i := 0; i < len(rows); i += batchSize { // to avoid unnessary load in database
		end := i + batchSize
		if end > len(rows) {
			end = len(rows)
		}

		// Process current batch
		for _, row := range rows[i:end] {
			customerObjID, exists := customerMap[row.CustomerID]
			if !exists {
				customer := model.Customer{
					ID:         primitive.NewObjectID(),
					CustomerID: row.CustomerID,
					Name:       row.CustomerName,
					Email:      row.CustomerEmail,
					Address:    row.CustomerAddr,
					CreateAt:   time.Now(),
					UpdateAt:   time.Now(),
				}

				customerObjID = customer.ID
				customerMap[row.CustomerID] = customerObjID
				customers = append(customers, customer)
			}

			// --- Insert Product ---
			productObjID, exists := productMap[row.ProductID]
			if !exists {
				product := model.Product{
					ID:        primitive.NewObjectID(),
					ProductID: row.ProductID,
					Name:      row.ProductName,
					Category:  row.Category,
					CreateAt:  time.Now(),
					UpdateAt:  time.Now(),
				}

				productObjID = product.ID
				productMap[row.ProductID] = productObjID
				products = append(products, product)
			}

			// --- Parse Date ---
			saleDate, _ := time.Parse("02-01-2006", row.DateOfSale)

			// --- Insert Order ---
			totalAmount := float64(row.QuantitySold)*row.UnitPrice*(1.0-row.Discount) + row.ShippingCost
			order := model.Order{
				ID:            primitive.NewObjectID(),
				OrderID:       row.OrderID,
				CustomerID:    customerObjID,
				ProductID:     productObjID,
				Region:        row.Region,
				DateOfSale:    saleDate,
				QuantitySold:  row.QuantitySold,
				UnitPrice:     row.UnitPrice,
				Discount:      row.Discount,
				ShippingCost:  row.ShippingCost,
				PaymentMethod: row.PaymentMethod,
				TotalAmount:   totalAmount,
				CreateAt:      time.Now(),
				UpdateAt:      time.Now(),
			}
			orders = append(orders, order)
		}

		repo.BulkInsertCustomers(customers)
		repo.BulkInsertProducts(products)
		repo.BulkInsertOrders(orders)
	}
}
