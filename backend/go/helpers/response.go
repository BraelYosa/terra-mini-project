package helpers

import "github.com/labstack/echo/v4"

func ErrorResponse(c echo.Context, errorType int, data any, errorMessage ...string) error {

	var errorResponse map[string]any
	var errorResponseMessage string

	if len(errorMessage) > 0 {
		errorResponseMessage = errorMessage[0]
	} else {
		switch errorType {
		case 200:
			errorResponseMessage = "OK"
		case 201:
			errorResponseMessage = "CREATED"
		case 203:
			errorResponseMessage = "FORBIDDEN"
		case 400:
			errorResponseMessage = "BAD_REQUEST"
		case 404:
			errorResponseMessage = "DATA_NOT_FOUND"
		case 500:
			errorResponseMessage = "INTERNAL_SERVER_ERROR"
		}
	}

	errorResponse = map[string]any{
		"Error_Type": errorType,
		"Message":    errorResponseMessage,
		"Data":       data,
	}

	return c.JSON(errorType, errorResponse)

}
