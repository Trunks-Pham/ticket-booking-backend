package models

import (
	"gorm.io/gorm"
)

// UserRole đại diện cho vai trò người dùng trong hệ thống.
type UserRole string

const (
	Manager  UserRole = "manager"  // Quản lý
	Customer UserRole = "customer" // Khách hàng
)

// User đại diện cho thông tin người dùng.
// @Description Đối tượng User chứa các thông tin cơ bản của người dùng như họ tên, email, mật khẩu, số điện thoại, số căn cước, hộ chiếu, vai trò và trạng thái.
// @Param firstName body string true "Tên của người dùng"
// @Param lastName body string true "Họ của người dùng"
// @Param email body string true "Email của người dùng"
// @Param password body string true "Mật khẩu người dùng"
// @Param phoneNumber body string true "Số điện thoại người dùng"
// @Param identityId body string false "Số căn cước của người dùng"
// @Param passport body string false "Số hộ chiếu của người dùng"
// @Param role body string true "Vai trò của người dùng (manager hoặc customer)"
// @Param status body boolean true "Trạng thái của người dùng (true: Hoạt động, false: Không hoạt động)"
type User struct {
	gorm.Model
	FirstName      string           `json:"firstName" gorm:"varchar(255);not null"`                                           // Tên của người dùng
	LastName       string           `json:"lastName" gorm:"varchar(255);not null"`                                            // Họ của người dùng
	Email          string           `json:"email" gorm:"varchar(255);not null"`                                               // Email của người dùng
	Password       string           `json:"password" gorm:"type:varchar(255);not null"`                                       // Mật khẩu người dùng
	PhoneNumber    string           `json:"phoneNumber" gorm:"varchar(255);not null"`                                         // Số điện thoại người dùng
	IdentityID     string           `json:"identityId" gorm:"varchar(255);"`                                                  // Số căn cước của người dùng
	Passport       string           `json:"passport" gorm:"varchar(256);"`                                                    // Số hộ chiếu của người dùng
	Role           UserRole         `json:"role" gorm:"text;default:customer"`                                                // Vai trò của người dùng (manager hoặc customer)
	Status         bool             `json:"status" gorm:"default:true"`                                                       // Trạng thái của người dùng (true: Hoạt động, false: Không hoạt động)
	BookingHistory []BookingHistory `json:"bookings" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` // Lịch sử đặt vé của người dùng
}

// AfterCreate kiểm tra và gán vai trò "manager" cho người dùng đầu tiên.
// @Summary Kiểm tra vai trò của người dùng sau khi tạo.
// @Description Sau khi tạo, nếu ID của người dùng là 1, vai trò sẽ được cập nhật thành "manager".
func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID == 1 {
		db.Model(u).Update("role", Manager)
	}
	return
}
