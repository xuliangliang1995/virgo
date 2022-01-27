package service

import "virgo/internal/store"

type Service interface {
	Users() UserService
}

type service struct {
	store store.Factory
}

func NewService(store store.Factory) Service  {
	return &service{
		store: store,
	}
}

func (s *service) Users() UserService {
	return newUserService(s)
}
