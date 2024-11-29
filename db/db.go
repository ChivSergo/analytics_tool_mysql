package db

import (
	"analytics_tool/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Init инициализирует подключение к базе данных и мигрирует модели
func Init() {
	dsn := "root:pass@tcp(localhost:3336)/mdb?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Проверка подключения к базе данных
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL database: %v", err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Миграция моделей
	err = DB.AutoMigrate(&models.Position{}, &models.Employee{}, &models.Task{}, &models.Timesheet{})
	if err != nil {
		log.Fatalf("Failed to migrate models: %v", err)
	}
}
