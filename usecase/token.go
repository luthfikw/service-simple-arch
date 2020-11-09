package usecase

import "github.com/koinworks/asgard-heimdal/libs/serror"

type TokenUsecase struct {
	userCase *UserUsecase
}

func NewTokenUsecase(userCase *UserUsecase) *TokenUsecase {
	return &TokenUsecase{
		userCase: userCase,
	}
}

func (ox *TokenUsecase) CreateToken() (resp string, errx serror.SError) {
	// TODO: do cross usecase stuff here

	return resp, errx
}
