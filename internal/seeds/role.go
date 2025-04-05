package seeds

import (
	"backend/internal/models"
	"context"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedDefaultRoles(db *gorm.DB) {
	ctx := context.Background()

	// Проверяем наличие роли admin
	var count int64
	if err := db.WithContext(ctx).Model(&models.Role{}).Where("name = ?", "admin").Count(&count).Error; err != nil {
		log.Fatalf("failed to check existing roles: %v", err)
	}

	if count == 0 {
		adminRole := models.Role{
			ID:   uuid.New(),
			Name: "admin",
			Desc: "Это роль администратора системы. Администратор имеет полный доступ ко всем функциям системы. Изменить и удалить эту роль нельзя.",
		}

		if err := db.WithContext(ctx).Create(&adminRole).Error; err != nil {
			log.Fatalf("failed to seed admin role: %v", err)
		}

		log.Println("✔ Seeded role: admin")
	}
}
