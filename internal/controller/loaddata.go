package controller

import (
	"lumel/internal/dataloader"
	"lumel/internal/response"
	"lumel/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

func LoadData(c *fiber.Ctx) error {
	logger.Log.Info("Data start loading in run in background")

	go dataloader.LoadSalesData()

	return response.SendSuccess(c, fiber.StatusOK, "Data start loading in run in background", nil)
}
