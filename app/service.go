package app

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/luthfikw/service-simple-arch/delivery/bivrost"
	"github.com/luthfikw/service-simple-arch/usecase"
)

func (ox *App) initService() serror.SError {
	userCase := usecase.NewUserUsecase(ox.DB)
	tokenCase := usecase.NewTokenUsecase(userCase)

	bivrost.NewBivrostHandler(ox.Bivrost, userCase, tokenCase)

	return nil
}
