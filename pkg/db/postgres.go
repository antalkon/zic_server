package db

import (
	"backend/internal/models"
	"backend/internal/seeds"
	"backend/pkg/config"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(cfg *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	log.Println("Connected to PostgreSQL successfully")

	// Автоматические миграции
	if err := db.AutoMigrate(
		&models.Room{},
		&models.Computer{},
		&models.Role{},

		&models.User{},
		&models.RefreshToken{},
	); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return nil, err
	}
	seeds.SeedDefaultRoles(db)

	return &Database{DB: db}, nil
}
