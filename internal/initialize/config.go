package initialize

import (
	"fmt"
	"github.com/Trunks-Pham/ticket-booking-backend/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viperConfig := viper.New()

	viperConfig.AddConfigPath("./config")
	viperConfig.SetConfigName("local")
	viperConfig.SetConfigType("yaml")

	err := viperConfig.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	if err := viperConfig.Unmarshal(&global.Config); err != nil {
		fmt.Println("Unable to decode configuration %v", err)
	}
}
