package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"newtestpostgre/entities"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {

	t.Run("ErrorGetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})
		userController.Get()(context)

		var response GetUsersResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, "Adlan", response.Data[0].Nama)
		assert.Equal(t, http.StatusOK, response.Code)
		//
	})
	t.Run("ErrorGetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users")

		falseUserController := New(MockFalseUserRepository{})
		falseUserController.Get()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "There is some error on server")
	})

}

func TestGetById(t *testing.T) {
	t.Run("GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.GetById()(context)

		response := GetUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)

	})
	t.Run("ErorGetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		falseUserController := New(MockFalseUserRepository{})
		falseUserController.GetById()(context)

		var response GetUserResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, response.Message, "not found")
	})

}

func TestUserRegister(t *testing.T) {
	t.Run("UserRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": "adlan123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})

		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		// assert.Equal(t, 201, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)
		assert.Equal(t, http.StatusCreated, response.Code)

	})
	t.Run("ErorUserRegister", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(&MockFalseUserRepository{})
		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("UserRegisterBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users")

		userController := New(MockUserRepository{})
		userController.UserRegister()(context)

		response := RegisterUserResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})

}

func TestUpdate(t *testing.T) {
	t.Run("Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": "adlan123",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Adlan", response.Data.Nama)

	})

	t.Run("ErrorUpdate", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodPut, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("UpdateBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Adlan",
			"email":    "adlan@adlan.com",
			"password": 123,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Update()(context)

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)

	})
}

func TestDelete(t *testing.T) {
	t.Run("DeleteUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockUserRepository{})
		userController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, nil, response.Data)

	})

	t.Run("ErrorDeleteUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		context := e.NewContext(req, res)
		context.SetPath("/users/:id")

		userController := New(&MockFalseUserRepository{})
		userController.Delete()(context)

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
}

// func TestLogin(t *testing.T) {
// 	t.Run("UserLogin", func(t *testing.T) {
// 		e := echo.New()

// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"email":    "adlan@adlan.com",
// 			"password": "adlan123",
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		userController := New(MockUserRepository{})
// 		userController.Login()(context)

// 		response := UserLoginResponseFormat{}
// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)
// 		fmt.Println(response)

// 		assert.Equal(t, 200, response.Code)
// 		assert.Equal(t, "adlan@adlan.com", response.Data.Email)

// 	})

// 	t.Run("ErrorLogin", func(t *testing.T) {
// 		e := echo.New()
// 		req := httptest.NewRequest(http.MethodPost, "/", nil)
// 		res := httptest.NewRecorder()
// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		userController := New(&MockFalseUserRepository{})
// 		userController.Login()(context)

// 		response := UserLoginResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)
// 		assert.Equal(t, "There is some problem from input", response.Message)

// 	})

// 	t.Run("UserLoginBind", func(t *testing.T) {
// 		e := echo.New()

// 		requestBody, _ := json.Marshal(map[string]interface{}{
// 			"email":    "adlan@adlan.com",
// 			"password": 123,
// 		})
// 		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
// 		res := httptest.NewRecorder()
// 		req.Header.Set("Content-Type", "application/json")

// 		context := e.NewContext(req, res)
// 		context.SetPath("/users/login")

// 		userController := New(MockUserRepository{})
// 		userController.Login()(context)

// 		response := UserLoginResponseFormat{}

// 		json.Unmarshal([]byte(res.Body.Bytes()), &response)

// 		assert.Equal(t, 400, response.Code)

// 	})
// }

type MockUserRepository struct{}

func (m MockUserRepository) Get() ([]entities.User, error) {
	return []entities.User{
		{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"},
	}, nil
}

func (m MockUserRepository) GetById(userId int) (entities.User, error) {
	return entities.User{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockUserRepository) UserRegister(newUser entities.User) (entities.User, error) {
	return entities.User{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

// func (m MockUserRepository) Login(data entities.User) (entities.User, error) {
// 	return entities.User{Email: "adlan@adlan.com", Password: "adlan123"}, nil
// }

func (m MockUserRepository) Update(userId int, newUser entities.User) (entities.User, error) {
	return entities.User{Nama: "Adlan", Email: "adlan@adlan.com", Password: "adlan123"}, nil
}

func (m MockUserRepository) Delete(userId int) error {
	return nil
}

type MockFalseUserRepository struct{}

func (m MockFalseUserRepository) Get() ([]entities.User, error) {
	return nil, errors.New("False User Object")
}
func (m MockFalseUserRepository) GetById(userId int) (entities.User, error) {
	return entities.User{}, errors.New("False Get Object")
}
func (m MockFalseUserRepository) UserRegister(newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Register Object")
}

// func (m MockFalseUserRepository) Login(data entities.User) (entities.User, error) {
// 	return entities.User{}, errors.New("False Login Object")
// }
func (m MockFalseUserRepository) Update(userId int, newUser entities.User) (entities.User, error) {
	return entities.User{}, errors.New("False Update Object")
}
func (m MockFalseUserRepository) Delete(userId int) error {
	return errors.New("False Delete Object")
}
