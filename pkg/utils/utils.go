/*
utils.go
Author: Naveenraj O M
Description: This utility package provides common functions used across the application,
such as environment loading, random number generation, authentication utilities,
and JWT token generation.
*/

package utils

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func ParseDateRange(c *fiber.Ctx) (time.Time, time.Time, error) {
	layout := "2006-01-02"
	startStr := c.Query("start")
	endStr := c.Query("end")

	startDate, err := time.Parse(layout, startStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid start date")
	}

	endDate, err := time.Parse(layout, endStr)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid end date")
	}

	return startDate, endDate, nil
}
