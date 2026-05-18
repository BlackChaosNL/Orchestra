package repositories

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func CreateUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser() (*User, err) {
	
}