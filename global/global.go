package global

import (
	"github.com/Trunks-Pham/ticket-booking-backend/pkg/settings"
	"gorm.io/gorm"
)

var (
	Config settings.Config
	Pdb    *gorm.DB
)
