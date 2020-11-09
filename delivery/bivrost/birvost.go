package bivrost

import (
	bv "github.com/koinworks/asgard-bivrost/service"
	"github.com/luthfikw/service-simple-arch/usecase"
)

type bvObject struct {
	Service   *bv.Service
	UserCase  *usecase.UserUsecase
	TokenCase *usecase.TokenUsecase
}

func NewBivrostHandler(
	svc *bv.Service,
	userCase *usecase.UserUsecase,
	tokenCase *usecase.TokenUsecase,
) {
	obj := &bvObject{
		Service:   svc,
		UserCase:  userCase,
		TokenCase: tokenCase,
	}

	obj.initRoute()
}
