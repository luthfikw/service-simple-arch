package app

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"

	"github.com/luthfikw/service-simple-arch/controllers"
	"github.com/luthfikw/service-simple-arch/deliveries/bivrost"
	"github.com/luthfikw/service-simple-arch/models"
)

func (ox *App) initService() serror.SError {
	models.InitModels(ox.DB)

	userControl := controllers.NewUserController()
	tokenControl := controllers.NewTokenController(userControl)

	bivrost.NewBivrostHandler(ox.Bivrost, userControl, tokenControl)

	return nil
}
