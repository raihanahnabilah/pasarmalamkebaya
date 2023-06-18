package usecase

import (
	"pasarmalamkebaya/entity"
	"pasarmalamkebaya/repository"
)

type OauthUsecase interface {
	FindUserByClientIDAndSecretID(clientID string, clientSecret string) (entity.OauthClient, error)
}

type oauthUsecase struct {
	oauthRepo repository.OauthRepo
}

func NewOauthUsecase(oauthRepo repository.OauthRepo) OauthUsecase {
	return &oauthUsecase{oauthRepo}
}

func (u *oauthUsecase) FindUserByClientIDAndSecretID(clientID string, clientSecret string) (entity.OauthClient, error) {
	return u.oauthRepo.FindUserByClientIDAndSecretID(clientID, clientSecret)
}
