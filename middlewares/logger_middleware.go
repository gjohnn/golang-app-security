package middlewares

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// LoggerMiddleware agrega logs a las solicitudes HTTP
func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.Logger()
}
