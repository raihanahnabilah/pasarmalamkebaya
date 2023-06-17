package repository

import (
	"pasarmalamkebaya/entity"

	"gorm.io/gorm"
)

// This is like the list of methods we're gonna define in this repo!
type UserRepo interface {
	FindUserByEmail(email string) (entity.User, error) // To find the user!
	CreateUser(user entity.User) (entity.User, error)  // To create the user!
}

// This is the list of things we inject here!
type userRepo struct {
	db *gorm.DB
}

// This is like a constructor!
func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

// The FindUserById method
func (r *userRepo) FindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// The CreateUser method
func (r *userRepo) CreateUser(user entity.User) (entity.User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
