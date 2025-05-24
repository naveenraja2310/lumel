// router.go
package router

import (
	"lumel/internal/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	logs "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// GetRouter initializes and returns the main Fiber application with configured routes.
func GetRouter() *fiber.App {
	app := fiber.New() // Initialize a new Fiber app

	//adding cors
	app.Use(cors.New())
	app.Use(logs.New()) // Add a custom logger middleware
	app.Use(recover.New())

	loadData := app.Group("/load-data")
	loadData.Get("/", controller.LoadData)

	analytics := app.Group("/analytics")
	analytics.Get("total-revenue", controller.GetTotalRevenue)
	analytics.Get("total-revenue-by-category", controller.GetRevenueByCategory)
	analytics.Get("total-revenue-by-product", controller.GetRevenueByProduct)
	analytics.Get("total-revenue-by-region", controller.GetRevenueByRegion)

	return app
}
