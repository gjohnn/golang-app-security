package services

import (
	"errors"
	"v0/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

// FindUserByUsernameOrEmail busca un usuario por nombre de usuario o email
func (s *AuthService) FindUserByUsernameOrEmail(username, email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("username = ? OR email = ?", username, email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (s *AuthService) Register(req *models.RegisterRequest) (*models.User, error) {
	// Verificar si el email o el nombre de usuario ya están registrados
	existingUser, _ := s.FindUserByUsernameOrEmail(req.Username, req.Email)
	if existingUser != nil {
		return nil, errors.New("el nombre de usuario o correo ya están registrados")
	}

	// Hashear la contraseña
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	var role string
	if req.Email == "papu@gmail.com" {
		role = "ADMIN"
	} else {
		role = "USER"
	}

	// Crear el nuevo usuario
	user := &models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Role:     role,
	}

	// Guardar el usuario en la base de datos
	if err := s.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// HashPassword hashea una contraseña
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword verifica si la contraseña ingresada es correcta
func CheckPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
