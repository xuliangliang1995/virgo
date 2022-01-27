package service

import (
	"virgo/internal/model"
	"virgo/internal/store"
)

type UserService interface {
	Create(user *model.User) error
	Update(user *model.User) error
	GetUserById(userId int64) (*model.User, error)
}

type userSrv struct {
	store store.Factory
}

func newUserService(srv *service) UserService {
	return &userSrv{
		store: srv.store,
	}
}

func (srv *userSrv) Create(user *model.User) error  {
	return srv.store.Users().Create(user)
}

func (srv *userSrv) Update(user *model.User) error {
	return srv.store.Users().Update(user)
}

func (srv *userSrv) GetUserById(userId int64) (*model.User, error) {
	return srv.store.Users().GetUserById(userId)
}