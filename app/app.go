package app

import (
	"github.com/jmoiron/sqlx"
	bv "github.com/koinworks/asgard-bivrost/service"
	"github.com/koinworks/asgard-heimdal/libs/serror"
)

type App struct {
	DB      *sqlx.DB
	Bivrost *bv.Service
}

func NewApp() *App {
	app := &App{}
	app.initBivrost()
	app.initService()

	return app
}

func (ox *App) Start() (errx serror.SError) {
	return errx
}
