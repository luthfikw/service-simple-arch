package app

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/luthfikw/u2"

	"github.com/koinworks/asgard-heimdal/constants/cservice"
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/koinworks/asgard-heimdal/utils/utint"
	"github.com/koinworks/asgard-heimdal/utils/utstring"
	"github.com/koinworks/asgard-heimdal/utils/uttime"

	"github.com/luthfikw/service-simple-arch/models"
)

func (ox *App) initDB() serror.SError {
	sqlConn := utstring.Env(cservice.DBConnStr, `
		host=__host__
		user=__user__
		password=__pwd__
		dbname=__name__
		sslmode=__sslMode__
		application_name=__appKey__
	`)
	sqlConn = u2.Binding(sqlConn, map[string]string{
		"host":    utstring.Env(cservice.DBHost, "127.0.0.1"),
		"user":    utstring.Env(cservice.DBUser, "root"),
		"pwd":     utstring.Env(cservice.DBPwd, ""),
		"name":    utstring.Env(cservice.DBName, "asgard"),
		"sslMode": utstring.Env(cservice.DBSSLMode, "disable"),
		"appKey":  utstring.Env(cservice.AppKey, models.AppKey),
		"appName": utstring.Env(cservice.AppName, models.AppName),
		"date":    uttime.ToString(uttime.DefaultDateFormat, time.Now()),
	})

	db, err := sqlx.Connect(utstring.Env(cservice.DBEngine, "postgres"), sqlConn)
	if err != nil {
		return serror.NewFromErrorc(err, fmt.Sprintf("failed connect to database %s", utstring.Env(cservice.DBName, "asgard")))
	}

	db.SetConnMaxLifetime(time.Minute * time.Duration(utint.StringToInt(utstring.Env(cservice.DBConnLifetime, "15"), 15)))
	db.SetMaxIdleConns(int(utint.StringToInt(utstring.Env(cservice.DBConnMaxIdle, "5"), 5)))
	db.SetMaxOpenConns(int(utint.StringToInt(utstring.Env(cservice.DBConnMaxOpen, "0"), 0)))

	ox.DB = db

	return nil
}
