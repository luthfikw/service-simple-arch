package controllers

import (
	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/luthfikw/service-simple-arch/models"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (ox *UserController) InsertUser(args models.InsertUser) (id int64, errx serror.SError) {
	id, errx = args.Execute(nil)
	if errx != nil {
		errx.AddComments("while execute insert user")
		return id, errx
	}

	return id, errx
}

func (ox *UserController) GetUserByID(id int64) (row models.User, errx serror.SError) {
	row = models.User{}

	errx = row.GetUserByID(id)
	if errx != nil {
		errx.AddComments("while get user by id")
		return row, errx
	}

	return row, errx
}

func (ox *UserController) GetUserList(pn models.Pagination) (rows []models.User, errx serror.SError) {
	rows, errx = models.GetUserList(pn)
	if errx != nil {
		errx.AddComments("while get user list")
		return rows, errx
	}

	return rows, errx
}
