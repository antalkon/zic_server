package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/models"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	if db == nil {
		panic("Database connection is nil in repository")
	}
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(client *models.User) error {
	return r.db.Table("users").Create(client).Error
}

func (r *AuthRepository) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	var client models.User
	err := r.db.Table("users").
		Where("phone_number = ? OR phone_number = ?", phoneNumber, phoneNumber[1:]).
		First(&client).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &client, nil
}

func (r *AuthRepository) CheckUsersCount() (int64, error) {
	var count int64
	err := r.db.Model(&models.User{}).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *AuthRepository) GetAdminRoleId() (uuid.UUID, error) {
	var role models.Role
	err := r.db.Table("roles").Where("name = ?", "admin").First(&role).Error
	if err != nil {
		return uuid.Nil, err
	}
	return role.ID, nil
}
