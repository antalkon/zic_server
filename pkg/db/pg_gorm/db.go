package pggorm

import (
	"fmt"
	"log"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// migrateModels выполняет миграцию моделей
func migrateModels() error {
	modelsToMigrate := []interface{}{
		&models.Computer{},
		&models.Room{},
		&models.CPULoad{},
		&models.RAMLoad{},
		&models.NetworkLoad{},
		&models.Class{},
		&models.PublicTasks{},
		&models.DeliveredTasks{},
	}

	for _, model := range modelsToMigrate {
		log.Printf("Migrating model: %T", model)
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Migration failed for model: %T, error: %v", model, err)
			return fmt.Errorf("failed to migrate model: %v", err)
		}
	}
	return nil
}

// InitDB инициализирует подключение к базе данных
func InitDB() error {
	// Загружаем конфигурацию из .env
	if err := config.LoadEnvConfig(); err != nil {
		log.Printf("Failed to load env config: %v", err)
		return fmt.Errorf("failed to load env config: %v", err)
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetEnvValue("DB_HOST"),     // Используйте "DB_HOST"
		config.GetEnvValue("DB_USER"),     // Используйте "DB_USER"
		config.GetEnvValue("DB_PASSWORD"), // Используйте "DB_PASSWORD"
		config.GetEnvValue("DB_NAME"),     // Убедитесь, что "DB_NAME" тоже используется
		config.GetEnvValue("DB_PORT"),     // Используйте "DB_PORT"
		config.GetEnvValue("DB_SSLMODE"),  // Используйте "DB_SSLMODE"
	)

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return fmt.Errorf("failed to connect database: %v", err)
	}

	DB = db

	log.Println("Connected to database successfully")

	// Выполняем миграцию моделей
	if err := migrateModels(); err != nil {
		log.Printf("Error during migration: %v", err)
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully")
	return nil
}
