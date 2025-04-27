package middleware

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		code    = http.StatusInternalServerError
		message interface{}
	)

	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		message = he.Message
	} else if ve, ok := err.(validator.ValidationErrors); ok {
		code = http.StatusBadRequest
		var errs []string
		for _, fe := range ve {
			field := strings.ToLower(fe.Field())
			switch fe.Tag() {
			case "required":
				errs = append(errs, field+" is required")
			case "len":
				errs = append(errs, field+" must be "+fe.Param()+" characters")
			case "numeric":
				errs = append(errs, field+" must be numeric")
			case "oneof":
				errs = append(errs, field+" must be one of "+fe.Param())
			case "gt":
				errs = append(errs, field+" must be greater than "+fe.Param())
			case "gte":
				errs = append(errs, field+" must be greater than or equal to "+fe.Param())
			case "url":
				errs = append(errs, field+" must be a valid URL")
			default:
				errs = append(errs, field+" is not valid")
			}
		}
		message = strings.Join(errs, ", ")
	} else {
		message = err.Error()
	}

	if !c.Response().Committed {
		c.JSON(code, echo.Map{
			"error": message,
		})
	}
}
