package usecase

import (
	"pasarmalamkebaya/entity"
	"pasarmalamkebaya/repository"
)

type OauthUsecase interface {
	FindUserByClientIDAndSecretID(clientID string, clientSecret string) (entity.OauthClient, error)
}

type oauthUsecase struct {
	oauthRepo repository.OauthClientRepo
}

func NewOauthUsecase(oauthRepo repository.OauthClientRepo) OauthUsecase {
	return &oauthUsecase{oauthRepo}
}

func (u *oauthUsecase) FindUserByClientIDAndSecretID(clientID string, clientSecret string) (entity.OauthClient, error) {
	return u.oauthRepo.FindUserByClientIDAndSecretID(clientID, clientSecret)
}
