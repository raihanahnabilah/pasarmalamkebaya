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
	mailUsecase MailUsecase
}

func NewRegisterUsecase(userRepo repository.UserRepo, userUsecase UserUsecase, mailUsecase MailUsecase) RegisterUsecase {
	return &registerUsecase{userRepo, userUsecase, mailUsecase}
}

func (u *registerUsecase) Register(input dto.RegisterRequestBody) (entity.User, error) {

	// When you register, you give inputs: Name, Email, Password

	// Check if the user is already registered
	userFound, err := u.userUsecase.FindUserByEmail(input.Email)
	// Have to check if the error is because the record is not found or any other errors?
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return entity.User{}, err
	}
	// Check if user already exists
	if userFound.UserID != 0 {
		return userFound, errors.New("The user already exists")
	}

	// Hashed the password!
	hashed, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	// Make a code!
	userRegistered := entity.User{
		Name:         input.Name,
		Email:        input.Email,
		Password:     string(hashed),
		CodeVerified: utils.RandStringBytes(32),
	}

	// Register
	user, err := u.userRepo.CreateUser(userRegistered)

	if err != nil {
		return entity.User{}, err
	}

	// Send verification to SendGrid
	sendVerification := dto.RegisterEmailVerification{
		Name:             input.Name,
		Subject:          "Verify your email to Pasar Malam Kebaya!",
		Email:            input.Email,
		VerificationCode: userRegistered.CodeVerified,
	}

	u.mailUsecase.SendEmailVerification(sendVerification)

	return user, nil

}
