package auth

import "newtestpostgre/entities"

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserLoginResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.User `json:"data"`
	Token   string        `json:"token"`
}
type FailedLoginResponseFormat struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// type LoginRequestFormat struct {
// 	Name string `json:"name" form:"name"`
// 	HP   string `json:"hp" form:"hp"`
// }

// type LoginResponseFormat struct {
// 	Message string        `json:"message"`
// 	Data    entities.User `json:"data"`
// 	Token   string        `json:"token"`
// }
