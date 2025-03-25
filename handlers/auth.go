package handlers

import (
	"net/http"
	"v0/models"
	"v0/services"
	"v0/utils"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	AuthService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req models.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	// Pasamos el request al servicio
	user, err := h.AuthService.Register(&req)
	if err != nil {
		return c.JSON(http.StatusConflict, map[string]string{"error": err.Error()})
	}

	// Generamos el token para el nuevo usuario
	token, _ := utils.GenerateToken(user.ID, user.Username, user.Email, user.Role)

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Usuario registrado",
		"token":   token,
	})
}
func (h *AuthHandler) Login(c echo.Context) error {
	var request models.LoginRequest

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Datos inválidos"})
	}

	user, err := h.AuthService.FindUserByUsernameOrEmail("", request.Email)
	if err != nil || user == nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Credenciales incorrectas"})
	}

	if err := services.CheckPassword(user.Password, request.Password); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Credenciales incorrectas"})
	}

	token, err := utils.GenerateToken(user.ID, user.Username, user.Email, user.Role)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Error generando token"})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login exitoso",
		"token":   token,
	})
}
