package store

import (
	"virgo/internal/model"
)


type UserStore interface {
	Create(user *model.User) error
	Update(user *model.User) error
}
