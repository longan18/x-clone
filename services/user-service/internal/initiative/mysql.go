package initiative

import (
	"fmt"
	"user-service/global"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySQL() {
	config := global.Config.Mysql
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.Dbname,
	)

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}
