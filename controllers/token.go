package controllers

import "github.com/koinworks/asgard-heimdal/libs/serror"

type TokenController struct {
	userControl *UserController
}

func NewTokenController(userControl *UserController) *TokenController {
	return &TokenController{
		userControl: userControl,
	}
}

func (ox *TokenController) CreateToken() (resp string, errx serror.SError) {
	// TODO: do cross controller stuff here

	return resp, errx
}
