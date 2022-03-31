package user

import "newtestpostgre/entities"

type User interface {
	Get() ([]entities.User, error)
	GetById(userId int) (entities.User, error)
	UserRegister(newUser entities.User) (entities.User, error)
	Update(userId int, newUser entities.User) (entities.User, error)
	Delete(userId int) error
}
