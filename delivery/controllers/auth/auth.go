package auth

import (
	"net/http"
	"newtestpostgre/configs"
	"newtestpostgre/delivery/controllers/common"
	"newtestpostgre/entities"
	"newtestpostgre/repository/auth"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (a AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		loginFormat := LoginRequest{}

		if err := c.Bind(&loginFormat); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		checkedUser, err := a.repo.Login(loginFormat.Email, loginFormat.Password)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, FailedLoginResponseFormat{
				Code:    http.StatusInternalServerError,
				Message: "Login Failed",
			})
		}

		token, err := GenerateToken(checkedUser)

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, map[string]interface{}{
				"Code":    http.StatusNotAcceptable,
				"message": "cannot process obtained value",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Code":    http.StatusOK,
			"message": "success login",
			"data":    checkedUser,
			"token":   token,
		})
	}
}

// func createToken(hp string) (string, error) {
// 	claims := jwt.MapClaims{}
// 	claims["authorized"] = true
// 	claims["id"] = 1
// 	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() //Token expires after 1 hour
// 	claims["name"] = hp

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	return token.SignedString([]byte("RAHASIA"))
// }

func GenerateToken(user entities.User) (string, error) {
	datas := jwt.MapClaims{}
	datas["id"] = user.ID
	datas["exp"] = time.Now().Add(time.Hour * 1).Unix() //1jam
	datas["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datas)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func ExtractTokenUserId(e echo.Context) float64 {

	user := e.Get("user").(*jwt.Token)
	if user.Valid {
		datas := user.Claims.(jwt.MapClaims)
		userId := datas["id"].(float64)
		return userId
	}

	return 0
}
