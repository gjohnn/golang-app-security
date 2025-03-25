package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PublicRoute permite acceso sin autenticación
func PublicRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Ruta pública"})
}

// UserRoute requiere autenticación
func UserRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Ruta protegida para usuarios"})
}

// AdminRoute requiere rol de ADMIN
func AdminRoute(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Ruta protegida para administradores"})
}
