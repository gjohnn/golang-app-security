package middlewares

import (
	"net/http"
	"strings"
	"v0/utils"

	"github.com/labstack/echo/v4"
)

// AuthMiddleware protege rutas con autenticación JWT y verifica el rol
func AuthMiddleware(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 1. Extraer el token del header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Token de autorización requerido",
					"details": "El header 'Authorization' está vacío",
				})
			}

			// 2. Verificar formato del token (Bearer token)
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Formato de token inválido",
					"details": "Formato esperado: 'Bearer <token>'",
				})
			}
			token := tokenParts[1]

			// 3. Validar token JWT
			claims, err := utils.ValidateToken(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]interface{}{
					"error":   "Token inválido",
					"details": err.Error(),
				})
			}

			// 4. Verificar que el token contiene el "sub" (usuario ID)
			userID, ok := claims["id"].(float64) // ID de usuario suele ser float64 en el token
			if !ok || userID == 0 {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Token no contiene identificador de usuario",
				})
			}
			// 5. Verificación de roles (si se especificaron)
			if len(allowedRoles) > 0 {
				userRole, ok := claims["role"].(string)
				if !ok || userRole == "" {
					return c.JSON(http.StatusForbidden, map[string]string{
						"error": "El token no contiene información de rol",
					})
				}

				if !containsRole(userRole, allowedRoles) {
					return c.JSON(http.StatusForbidden, map[string]interface{}{
						"error":   "Acceso denegado",
						"details": "Rol insuficiente. Se requiere uno de: " + strings.Join(allowedRoles, ", "),
					})
				}
			}

			// 6. Almacenar claims en contexto para handlers posteriores
			c.Set("user", claims)

			// 7. Continuar con la cadena de middlewares/handlers
			return next(c)
		}
	}
}

// containsRole verifica si el rol del usuario está en la lista de roles permitidos
func containsRole(userRole string, allowedRoles []string) bool {
	for _, role := range allowedRoles {
		if strings.EqualFold(userRole, role) {
			return true
		}
	}
	return false
}
