package pggorm

import (
	"fmt"
	"log"

	"github.com/antalkon/zic_server/internal/models"
	"github.com/antalkon/zic_server/pkg/config"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func migrateModels() error {
	modelsToMigrate := []interface{}{
		&models.Computer{},
		&models.Room{},
	}

	for _, model := range modelsToMigrate {
		log.Printf("Migrating model: %T", model) // Логируем каждую модель перед миграцией
		if err := DB.AutoMigrate(model); err != nil {
			log.Printf("Migration failed for model: %T, error: %v", model, err) // Логируем ошибки миграции
			return fmt.Errorf("failed to migrate model: %v", err)
		}
	}
	return nil
}

// InitDB - функция для инициализации подключения к базе данных
func InitDB() error {
	// Загружаем конфиг из файла
	if err := config.LoadDBConfig(); err != nil {
		log.Printf("Failed to load DB config: %v", err) // Логируем ошибки загрузки конфига
		return fmt.Errorf("failed to load DB config: %v", err)
	}

	// Формируем строку подключения
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		viper.GetString("host"),
		viper.GetString("user"),
		viper.GetString("password"),
		viper.GetString("dbname"),
		viper.GetString("port"),
		viper.GetString("sslmode"),
	)

	// Подключаемся к базе данных
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err) // Логируем ошибки подключения
		return fmt.Errorf("failed to connect database: %v", err)
	}

	// Сохраняем подключение в глобальную переменную
	DB = db

	log.Println("Connected to database successfully") // Логируем успешное подключение

	// Проверка работы миграций
	if err := migrateModels(); err != nil {
		log.Printf("Error during migration: %v", err) // Логируем ошибки миграции
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database migration completed successfully") // Логируем успешную миграцию
	return nil
}
