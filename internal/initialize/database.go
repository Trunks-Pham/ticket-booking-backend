package initialize

import (
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/Trunks-Pham/ticket-booking-backend/internal/models"

	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPostgreSql() {
	m := global.Config.PostgreSql
	uri := fmt.Sprintf(`
		host=%s user=%s dbname=%s password=%s sslmode=%s port=%v`,
		m.Host, m.Username, m.Dbname, m.Password, m.Sslmode, m.Port,
	)

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	log.Info("Connected to the database")

	//if err := DBMigrator(db); err != nil {
	//	log.Fatalf("Unable to migrate: %v", err)
	//}

	global.Pdb = db
}

func DBMigrator(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Flight{},
		&models.Ticket{},
		&models.BookingHistory{},
	)
	return err
}
