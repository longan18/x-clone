package initiative

import (
	"auth-service/global"
	"fmt"

	"github.com/spf13/viper"
)

func InitLoadConfig() {
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err.Error()))
	}

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration: %s", err.Error())
	}
}
