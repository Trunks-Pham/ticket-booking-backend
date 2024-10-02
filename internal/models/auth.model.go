package models

import (
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

// RegisterCredentials chứa thông tin cần thiết để đăng ký người dùng.
// @Description Đối tượng RegisterCredentials chứa các trường cần thiết cho việc đăng ký tài khoản trong hệ thống.
// @Param email body string true "Địa chỉ email"
// @Param password body string true "Mật khẩu"
// @Param firstName body string true "Họ"
// @Param lastName body string true "Tên"
// @Param phoneNumber body string true "Số điện thoại"
// @Param identityId body string false "ID cá nhân (tuỳ chọn)"
// @Param passport body string false "Hộ chiếu (tuỳ chọn)"
type RegisterCredentials struct {
	Email       string `json:"email" validate:"required"`       // Email người dùng
	Password    string `json:"password" validate:"required"`    // Mật khẩu người dùng
	FirstName   string `json:"firstName" validate:"required"`   // Họ người dùng
	LastName    string `json:"lastName" validate:"required"`    // Tên người dùng
	PhoneNumber string `json:"phoneNumber" validate:"required"` // Số điện thoại người dùng
	IdentityID  string `json:"identityId" validate:"omitempty"` // ID cá nhân (tuỳ chọn)
	Passport    string `json:"passport" validate:"omitempty"`   // Hộ chiếu (tuỳ chọn)
}

// LoginCredentials chứa thông tin cần thiết để đăng nhập người dùng.
// @Description Đối tượng LoginCredentials chứa email và mật khẩu để đăng nhập người dùng vào hệ thống.
// @Param email body string true "Địa chỉ email"
// @Param password body string true "Mật khẩu"
type LoginCredentials struct {
	Email    string `json:"email" validate:"required"`    // Email người dùng
	Password string `json:"password" validate:"required"` // Mật khẩu người dùng
}

// MatchesHash kiểm tra xem mật khẩu có khớp với mã băm không.
// @Description Hàm này so sánh mật khẩu dạng text với mật khẩu đã được mã băm để xác thực xem chúng có khớp nhau không.
// @Param password string true "Mật khẩu dạng text"
// @Param hash string true "Mật khẩu đã mã băm"
// @Success 200 {boolean} bool "Trả về true nếu mật khẩu khớp, ngược lại trả về false"
func MatchesHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// IsValidEmail kiểm tra xem địa chỉ email có hợp lệ không.
// @Description Hàm này kiểm tra xem email cung cấp có phải là một địa chỉ email hợp lệ hay không.
// @Param email string true "Địa chỉ email"
// @Success 200 {boolean} bool "Trả về true nếu email hợp lệ, ngược lại trả về false"
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
