package repository

import (
	"pasarmalamkebaya/entity"

	"gorm.io/gorm"
)

// Client and Access Token -- Not sure what's really the difference, so I mixed them.
// Authentication Flow: Login -> get access token. To get access token -> confirm email exist (user usecase) -> confirm password (user usecase)
// --> confirm client ID and client secret exists (on oauth_client table) --> they exist? create an access token, push it to oauth_access_token (oauth_acces)

type OauthAccessRepo interface {
	CreateAccessToken(accessToken entity.OauthAccessToken) (entity.OauthAccessToken, error)
}

type oauthAccessRepo struct {
	db *gorm.DB
}

func NewOauthAccessRepository(db *gorm.DB) *oauthAccessRepo {
	return &oauthAccessRepo{db}
}

func (r *oauthAccessRepo) CreateAccessToken(accessToken entity.OauthAccessToken) (entity.OauthAccessToken, error) {
	err := r.db.Create(&accessToken).Error
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
