package user

import (
	"virgo/internal/service"
	"virgo/internal/store"
)

type UserController struct {
	srv service.Service
}

func NewUserController(store store.Factory) *UserController  {
	return &UserController{
		srv: service.NewService(store),
	}
}
