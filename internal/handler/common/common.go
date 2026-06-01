package common

import (
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Validate is the shared validator instance
var Validate = validator.New()

// IntQuery parses an integer query parameter with a default value
func IntQuery(c *fiber.Ctx, key string, defaultVal int) int {
	val := c.Query(key)
	if val == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		return defaultVal
	}
	return n
}

// GetUserID extracts user ID from JWT locals (set by middleware)
func GetUserID(c *fiber.Ctx) string {
	if id, ok := c.Locals("userID").(string); ok {
		return id
	}
	return ""
}

// FormatValidationErrors converts validator errors to a readable map
func FormatValidationErrors(err error) map[string]string {
	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["general"] = err.Error()
		return errors
	}

	for _, e := range validationErrors {
		field := e.Field()
		switch e.Tag() {
		case "required":
			errors[field] = field + " is required"
		case "min":
			errors[field] = field + " must be at least " + e.Param() + " characters"
		case "max":
			errors[field] = field + " must be at most " + e.Param() + " characters"
		case "gt":
			errors[field] = field + " must be greater than " + e.Param()
		case "gte":
			errors[field] = field + " must be greater than or equal to " + e.Param()
		case "uuid":
			errors[field] = field + " must be a valid UUID"
		case "url":
			errors[field] = field + " must be a valid URL"
		case "oneof":
			errors[field] = field + " must be one of: " + e.Param()
		default:
			errors[field] = field + " is invalid"
		}
	}

	return errors
}
