package usecase

import (
	"os"
	"pasarmalamkebaya/dto"
	"pasarmalamkebaya/entity"
	"pasarmalamkebaya/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase interface {
	Login(input dto.LoginRequestBody) (dto.LoginResponse, error)
}

type loginUsecase struct {
	oauthRepo    repository.OauthRepo
	oauthUsecase OauthUsecase
	userUsecase  UserUsecase
}

func NewLoginUsecase(oauthRepo repository.OauthRepo, oauthUsecase OauthUsecase, userUsecase UserUsecase) LoginUsecase {
	return &loginUsecase{oauthRepo, oauthUsecase, userUsecase}
}

// TODO: Please implement this!
// Put the LOGIN!
func (u *loginUsecase) Login(input dto.LoginRequestBody) (dto.LoginResponse, error) {

	// See if client_id and client_secret are registered to a database
	oauthClient, err := u.oauthUsecase.FindUserByClientIDAndSecretID(input.ClientID, input.ClientSecret)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// Find if the user exists via email -- call the user Usecase "find user by email"
	user, err := u.userUsecase.FindUserByEmail(input.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// Hashed the password and see if they are the same -- call the user Usecase, but get the password! "find user by email"
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// Create the JWT Token! -- call the Oauth Usecase "create"
	// JWT key
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	// get the expiration time -- 5 minutes only!
	expirationTime := time.Now().Add(365 * 24 * time.Hour)
	// create the JWT claims, which includes the email and expiry time!
	claims := &dto.Claims{
		Email: input.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			// in jwt, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	// Insert data to oauth access token table
	accessToken := entity.OauthAccessToken{
		ClientID:  &oauthClient.ID,
		UserID:    user.UserID,
		Token:     tokenString,
		Scope:     "*",
		ExpiredAt: &claims.ExpiresAt.Time,
	}
	oauthAccessToken, err := u.oauthRepo.CreateAccessToken(accessToken)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	type LoginResponse struct {
		Token     string `json:"access_token"`
		Type      string `json:"Bearer"`
		ExpiredAt string `json:"expired_at"`
		Scope     string `json:"scope"`
	}

	// Get the login response!
	loginResponse := dto.LoginResponse{
		Token:     oauthAccessToken.Token,
		Type:      "Bearer",
		ExpiredAt: expirationTime.Format(time.RFC3339),
		Scope:     "*",
	}

	return loginResponse, nil
}
