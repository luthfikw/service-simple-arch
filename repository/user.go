package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/luthfikw/service-simple-arch/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) (obj *UserRepository) {
	obj = &UserRepository{
		db: db,
	}
	return obj
}

func (ox *UserRepository) InsertUser(args models.InsertUserArguments) (id int64, errx serror.SError) {
	// TODO: do query to db here

	return id, errx
}

func (ox *UserRepository) UpdateUser(args models.UpdateUserArguments) (errx serror.SError) {
	// TODO: do query to db here

	return errx
}
