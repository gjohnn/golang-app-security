package main

import (
	"log"
	"net/http"
	"v0/database"
	"v0/handlers"
	"v0/middlewares"
	"v0/services"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	authService := services.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)

	e := echo.New()
	// Configurar CORS personalizado
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://my-frontend.com"},                 // Permite estos orígenes
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}, // Métodos permitidos
		AllowHeaders:     []string{"Content-Type", "Authorization"},                                    // Cabeceras permitidas
		AllowCredentials: true,                                                                         // Permite el envío de cookies y credenciales
	}))

	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middlewares.LoggerMiddleware())

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)

	// Rutas publicas
	public := e.Group("/public")
	public.GET("", handlers.PublicRoute)

	// Rutas de USER y tambien ADMIN
	protected := e.Group("/user", middlewares.AuthMiddleware("USER", "ADMIN"))
	protected.GET("", handlers.UserRoute)

	// Rutas de ADMIN
	admin := e.Group("/admin", middlewares.AuthMiddleware("ADMIN"))
	admin.GET("", handlers.AdminRoute)

	e.Start(":8080")
}
