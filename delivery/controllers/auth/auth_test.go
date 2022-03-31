package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"newtestpostgre/configs"
	"newtestpostgre/entities"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// jwtToken := ""
	t.Run("UserLogin", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "adlan@adlan.com",
			"password": "adlan123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		authController := New(MockAuthRepository{})
		authController.Login()(context)

		response := UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		fmt.Println(response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "adlan@adlan.com", response.Data.Email)

	})

	t.Run("ErrorLogin", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"email":    "adlan@adlan.com",
			"password": "ahaaa",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		authController := New(MockFalseAuthRepository{})
		authController.Login()(context)

		response := UserLoginResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "Login Failed", response.Message)

	})

	t.Run("UserLoginBind", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":     "testBind",
			"password": 123,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		context := e.NewContext(req, res)
		context.SetPath("/users/login")

		authController := New(MockAuthRepository{})
		authController.Login()(context)

		response := UserLoginResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
		assert.Equal(t, "There is some problem from input", response.Message)

	})
}

// type MockGenerateToken struct{}

// type MockExtractTokenUserId struct{}

type MockAuthRepository struct{}

func (m MockAuthRepository) Login(email, password string) (entities.User, error) {
	return entities.User{Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockAuthRepository) GenerateToken(data entities.User) (string, error) {
	datas := jwt.MapClaims{}
	datas["id"] = 1
	datas["exp"] = time.Now().Add(time.Hour * 1).Unix() //1jam
	datas["authorized"] = true
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, datas)
	return token.SignedString([]byte(configs.JWT_SECRET))
}

func (m MockAuthRepository) ExtractTokenUserId(e echo.Context) float64 {

	// user := e.Get("user").(*jwt.Token)
	// if user.Valid {
	// 	datas := user.Claims.(jwt.MapClaims)
	// 	userId := datas["id"].(float64)
	// 	return userId
	// }

	return 1
}

type MockFalseAuthRepository struct{}

func (m MockFalseAuthRepository) Login(email, password string) (entities.User, error) {
	return entities.User{}, errors.New("False Login Object")
}
