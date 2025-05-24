package response

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

var UserAlreadyExists = "User already exists"
var DomainNameAlreadyExists = "Domain name already exists"
var OrganizationNameAlreadyExists = "Organization name already exists"

// ErrorResponse defines the structure for API error responses
type ErrorResponse struct {
	ApiPath      string    `json:"apiPath"`
	ErrorCode    int       `json:"errorCode"`
	ErrorMessage string    `json:"errorMessage"`
	ErrorTime    time.Time `json:"errorTime"`
}

// SuccessResponse defines the structure for API success responses
type SuccessResponse struct {
	StatusCode    int         `json:"statusCode"`
	StatusMessage string      `json:"statusMessage"`
	Data          interface{} `json:"data"`
}

// CodeMessage Response
type CodeMessage struct {
	Error CustomError `json:"error"`
}

type CustomError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// SendError sends a standardized error response
func SendError(ctx *fiber.Ctx, statusCode int, message string, err error) error {
	return ctx.Status(statusCode).JSON(ErrorResponse{
		ApiPath:      ctx.OriginalURL(),
		ErrorCode:    statusCode,
		ErrorMessage: formatErrorMessage(message, err),
		ErrorTime:    time.Now(),
	})
}

// SendSuccess sends a standardized success response
func SendSuccess(ctx *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return ctx.Status(statusCode).JSON(SuccessResponse{
		StatusCode:    statusCode,
		StatusMessage: message,
		Data:          data,
	})
}

// formatErrorMessage formats the error message to include details if an error exists
func formatErrorMessage(message string, err error) string {
	if err != nil {
		return message + ": " + err.Error()
	}
	return message
}
