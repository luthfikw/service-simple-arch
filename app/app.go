package app

import (
	"github.com/jmoiron/sqlx"

	bv "github.com/koinworks/asgard-bivrost/service"
	"github.com/koinworks/asgard-heimdal/libs/logger"
	"github.com/koinworks/asgard-heimdal/libs/serror"
)

type App struct {
	DB      *sqlx.DB
	Bivrost *bv.Service
}

func NewApp() *App {
	app := &App{}

	most(app.initDB())
	most(app.initBivrost())
	most(app.initService())

	return app
}

func most(errx serror.SError) {
	if errx != nil {
		logger.Panic(errx)
	}
}

func (ox *App) Start() (errx serror.SError) {
	return errx
}
