package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/pkg/settings"

	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"
	"github.com/Trunks-Pham/ticket-booking-backend/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthService là một service chịu trách nhiệm cho các thao tác xác thực người dùng.
type AuthService struct {
	repository models.IAuthRepository // Repository để tương tác với cơ sở dữ liệu người dùng
	config     settings.Config        // Cấu hình ứng dụng
}

// Login cho phép người dùng đăng nhập bằng cách cung cấp email và mật khẩu.
// Nó trả về một JWT token và thông tin người dùng nếu đăng nhập thành công.
func (s *AuthService) Login(ctx context.Context, loginData *models.LoginCredentials) (string, *models.User, error) {
	user, err := s.repository.GetUser(ctx, "email = ?", loginData.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", nil, fmt.Errorf("invalid credentials")
		}
		return "", nil, err
	}

	if !models.MatchesHash(loginData.Password, user.Password) {
		return "", nil, fmt.Errorf("invalid credentials")
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, s.config.Authentication.JwtScretKey)

	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// Register cho phép người dùng đăng ký tài khoản mới.
// Nó trả về một JWT token và thông tin người dùng nếu đăng ký thành công.
func (s *AuthService) Register(ctx context.Context, registerData *models.RegisterCredentials) (string, *models.User, error) {
	if !models.IsValidEmail(registerData.Email) {
		return "", nil, fmt.Errorf("please, provide a valid email to register")
	}

	if registerData.IdentityID == "" && registerData.Passport == "" {
		return "", nil, fmt.Errorf("please, provide at least one of identityID or register")
	}

	if _, err := s.repository.GetUser(ctx, "email = ?", registerData.Email); !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", nil, fmt.Errorf("the user email is already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil, err
	}

	registerData.Password = string(hashedPassword)

	user, err := s.repository.RegisterUser(ctx, registerData)
	if err != nil {
		return "", nil, err
	}

	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 168).Unix(),
	}

	// Generate the JWT
	token, err := utils.GenerateJWT(claims, jwt.SigningMethodHS256, s.config.Authentication.JwtScretKey)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// NewAuthService khởi tạo một instance mới của AuthService với repository cho trước.
func NewAuthService(repository models.IAuthRepository) models.IAuthService {
	return &AuthService{
		repository: repository,
		config:     global.Config,
	}
}
