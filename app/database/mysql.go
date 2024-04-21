package database

import (
	"fmt"
	"lendra/app/config"
	pd "lendra/features/product/data"
	ud "lendra/features/user/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMySQL(cfg *config.AppConfig) *gorm.DB {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&ud.User{},
		&pd.Product{},
	)

	return DB
}
