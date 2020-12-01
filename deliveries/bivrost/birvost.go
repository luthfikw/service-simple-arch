package bivrost

import (
	bv "github.com/koinworks/asgard-bivrost/service"

	"github.com/luthfikw/service-simple-arch/controllers"
)

type bvObject struct {
	Service      *bv.Service
	UserControl  *controllers.UserController
	TokenControl *controllers.TokenController
}

func NewBivrostHandler(
	svc *bv.Service,
	userControl *controllers.UserController,
	tokenControl *controllers.TokenController,
) {
	obj := &bvObject{
		Service:      svc,
		UserControl:  userControl,
		TokenControl: tokenControl,
	}

	obj.initRoute()
}
