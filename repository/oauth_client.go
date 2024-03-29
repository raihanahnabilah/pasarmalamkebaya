package repository

import (
	"pasarmalamkebaya/entity"

	"gorm.io/gorm"
)

// Client and Access Token -- Not sure what's really the difference, so I mixed them.
// Authentication Flow: Login -> get access token. To get access token -> confirm email exist (user usecase) -> confirm password (user usecase)
// --> confirm client ID and client secret exists (on oauth_client table) --> they exist? create an access token, push it to oauth_access_token (oauth_acces)

type OauthClientRepo interface {
	FindUserByClientIDAndSecretID(clientID string, clientSecret string) (entity.OauthClient, error)
}

type oauthClientRepo struct {
	db *gorm.DB
}

func NewOauthClientRepository(db *gorm.DB) *oauthClientRepo {
	return &oauthClientRepo{db}
}

func (r *oauthClientRepo) FindUserByClientIDAndSecretID(clientID string, clientSecret string) (entity.OauthClient, error) {
	var oauth entity.OauthClient

	err := r.db.Where("client_id = ?", clientID).Where("client_secret = ?", clientSecret).First(&oauth).Error
	if err != nil {
		return oauth, err
	}

	return oauth, nil
}
