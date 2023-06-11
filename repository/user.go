package repository

import (
	e "pasarmalamkebaya/entity"

	"gorm.io/gorm"
)

// This is like the list of methods we're gonna define in this repo!
type UserRepo interface {
	FindById(id int) (e.User, error) // To find the user!
}

// This is the list of things we inject here!
type userRepo struct {
	db *gorm.DB
}

// This is like a constructor!
func NewUserRepository(db *gorm.DB) *userRepo {
	return &userRepo{db}
}

// The methods under the userRepo!
func (r *userRepo) FindById(id int) (e.User, error) {
	var user e.User

	err := r.db.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
