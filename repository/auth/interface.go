package auth

import "newtestpostgre/entities"

type Auth interface {
	Login(email, password string) (entities.User, error)
}
