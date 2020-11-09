package usecase

import (
	"github.com/jmoiron/sqlx"

	"github.com/koinworks/asgard-heimdal/libs/serror"
	"github.com/luthfikw/service-simple-arch/models"
	"github.com/luthfikw/service-simple-arch/repository"
)

type UserUsecase struct {
	db       *sqlx.DB
	userRepo *repository.UserRepository
}

func NewUserUsecase(db *sqlx.DB) *UserUsecase {
	return &UserUsecase{
		db:       db,
		userRepo: repository.NewUserRepository(db),
	}
}

func (ox *UserUsecase) InsertUser(args models.InsertUserArguments) (id int64, errx serror.SError) {
	return ox.userRepo.InsertUser(args)
}

func (ox *UserUsecase) UpdateUser(args models.UpdateUserArguments) (errx serror.SError) {
	return ox.userRepo.UpdateUser(args)
}
