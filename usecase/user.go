package usecase

import (
	"pasarmalamkebaya/entity"
	"pasarmalamkebaya/repository"
)

type UserUsecase interface {
	FindUserByEmail(email string) (entity.User, error)
}

type userUsecase struct {
	userRepo repository.UserRepo
}

func (u *userUsecase) FindUserByEmail(email string) (entity.User, error) {
	return u.userRepo.FindUserByEmail(email)

}

func NewUserUsecase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecase{userRepo}
}
