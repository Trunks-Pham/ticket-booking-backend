package initialize

import (
	"log"

	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/spf13/viper"
)

// LoadConfig loads configuration from the yaml file into global.Config
func LoadConfig() {
	viperConfig := viper.New()

	// Thiết lập đường dẫn và tên file cấu hình
	viperConfig.AddConfigPath("./config")
	viperConfig.SetConfigName("local") // Tên file không có đuôi mở rộng
	viperConfig.SetConfigType("yaml")  // Định dạng của file là YAML

	// Đọc file cấu hình
	err := viperConfig.ReadInConfig()
	if err != nil {
		log.Fatalf("❌ Lỗi khi đọc file cấu hình: %v \n", err)
	}

	// Giải mã cấu hình vào biến global.Config
	if err := viperConfig.Unmarshal(&global.Config); err != nil {
		log.Fatalf("❌ Không thể giải mã cấu hình: %v", err)
	}

	log.Println("✅ Cấu hình đã được tải thành công.")
}
