package dao

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"go_web/config"
	"go_web/models"
)

func InitDB() *gorm.DB {
	cfg := config.LoadConfig()

	connStr := fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
		cfg.Database.Password,
	)

	var err error
	DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// 启用详细日志
	//DB.LogMode(true)

	// 自动迁移
	DB.AutoMigrate(&models.Todo{})

	return DB
}
