package usecase

import (
	"errors"
	"pasarmalamkebaya/dto"
	"pasarmalamkebaya/entity"
	"pasarmalamkebaya/repository"
	"pasarmalamkebaya/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// The registration should have "Register" usercase
type RegisterUsecase interface {
	Register(dto dto.RegisterRequestBody) (entity.User, error) // To find the user!
}

// Plugins from other parts
type registerUsecase struct {
	userRepo    repository.UserRepo
	userUsecase UserUsecase
}

func NewRegisterUsecase(userRepo repository.UserRepo, userUsecase UserUsecase) RegisterUsecase {
	return &registerUsecase{userRepo, userUsecase}
}

func (u *registerUsecase) Register(dto dto.RegisterRequestBody) (entity.User, error) {

	// When you register, you give inputs: Name, Email, Password

	// Check if the user is already registered
	userFound, err := u.userUsecase.FindUserByEmail(dto.Email)
	// Have to check if the error is because the record is not found or any other errors?
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, err
	}
	// Check if user already exists
	if userFound.UserID != 0 {
		return userFound, errors.New("The user already exists")
	}

	// Hashed the password!
	hashed, _ := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)

	// Make a code!
	userRegistered := entity.User{
		Name:         dto.Name,
		Email:        dto.Email,
		Password:     string(hashed),
		CodeVerified: utils.RandStringBytes(32),
	}

	// Register
	user, err := u.userRepo.CreateUser(userRegistered)

	if err != nil {
		return entity.User{}, err
	}

	// Send to SENDGRID

	return user, nil

}
