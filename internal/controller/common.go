package controller

import (
	"lumel/internal/repo"

	"github.com/gofiber/fiber/v2"
)

type CommonController struct {
	repo *repo.CommonRepo
}

func NewCommonController(repo *repo.CommonRepo) *CommonController {
	return &CommonController{
		repo: repo,
	}
}

func (c *CommonController) HealthCheckHandler(ctx *fiber.Ctx) error {
	ctx.SendString("This lumel service is up!")
	return nil
}
