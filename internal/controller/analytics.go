package controller

import (
	"lumel/internal/repo"
	"lumel/internal/response"
	"lumel/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func GetTotalRevenue(c *fiber.Ctx) error {
	startDate, endDate, err := utils.ParseDateRange(c)
	if err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid date range", err)
	}

	totalRevenue, err := repo.GetTotalRevenue(startDate, endDate)
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Failed to calculate total revenue", err)
	}

	return response.SendSuccess(c, fiber.StatusOK, "Total revenue retrieved successfully", fiber.Map{"total_revenue": totalRevenue})
}

func GetRevenueByProduct(c *fiber.Ctx) error {
	startDate, endDate, err := utils.ParseDateRange(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	totalRevenue, err := repo.GroupedRevenue(startDate, endDate, "product_id")
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Failed to calculate total revenue", err)
	}

	return response.SendSuccess(c, fiber.StatusOK, "Total revenue by product", fiber.Map{"total_revenue": totalRevenue})
}

func GetRevenueByCategory(c *fiber.Ctx) error {

	startDate, endDate, err := utils.ParseDateRange(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	totalRevenue, err := repo.GroupedRevenue(startDate, endDate, "category")
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Failed to calculate total revenue", err)
	}
	return response.SendSuccess(c, fiber.StatusOK, "Total revenue by category", fiber.Map{"total_revenue": totalRevenue})
}

func GetRevenueByRegion(c *fiber.Ctx) error {

	startDate, endDate, err := utils.ParseDateRange(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	totalRevenue, err := repo.GroupedRevenue(startDate, endDate, "region")
	if err != nil {
		return response.SendError(c, fiber.StatusInternalServerError, "Failed to calculate total revenue", err)
	}
	return response.SendSuccess(c, fiber.StatusOK, "Total revenue by region", fiber.Map{"total_revenue": totalRevenue})
}
