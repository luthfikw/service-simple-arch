package bivrost

import (
	"net/http"

	bv "github.com/koinworks/asgard-bivrost/service"
)

func (ox *bvObject) welcome(ctx *bv.Context) bv.Result {
	return ctx.JSONResponse(http.StatusOK, bv.ResponseBody{
		Message: map[string]string{
			"en": "Welcome",
			"id": "Selamat datang",
		},
	})
}
