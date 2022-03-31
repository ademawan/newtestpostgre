package project

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"newtestpostgre/configs"
	"newtestpostgre/delivery/controllers/auth"
	"newtestpostgre/entities"
	"testing"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	jwtToken := ""
	t.Run("Get jwtToken", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("GET", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")

		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects")

		projectController := New(MockProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response GetProjectsResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Success Get All Project", response.Message)
		//
	})
	t.Run("error GET", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects")

		falseprojectController := New(MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(falseprojectController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}
		var response GetProjectsResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)
	})

}

func TestGetById(t *testing.T) {
	jwtToken := ""
	t.Run("Get jwtToken", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})
	t.Run("GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Success Get Project", response.Message)

	})
	t.Run("ErorGetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		falseprojectController := New(MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(falseprojectController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		var response GetProjectResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, http.StatusNotFound, response.Code)
		assert.Equal(t, "Not Found", response.Message)
	})

}

func TestProjectRegister(t *testing.T) {

	jwtToken := ""
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("ProjectRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "ProjectKu",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects/register")

		projectController := New(MockProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.ProjectRegister())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := RegisterProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusCreated, response.Code)

	})
	t.Run("ErorProjectRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"ed": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects/register")

		projectController := New(MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.ProjectRegister())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := RegisterProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("ProjectRegisterBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		fmt.Println(req)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects/register")

		projectController := New(MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.ProjectRegister())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := RegisterProjectResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "There is some problem from input", response.Message)

	})

}

func TestUpdate(t *testing.T) {

	jwtToken := ""
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})
	t.Run("Update", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "ProjectKu",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Success Update Project", response.Message)

	})

	t.Run("ErrorUpdate", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": "ProjectKu",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
	t.Run("error UpdateBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, "There is some problem from input", response.Message)

	})
}

func TestDelete(t *testing.T) {
	jwtToken := ""
	t.Run("Test Login", func(t *testing.T) {
		e := echo.New()

		requestBody, _ := json.Marshal(map[string]string{
			"email":    "test@gmail.com",
			"password": "xyz",
		})

		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})
	t.Run("Delete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "Success Delete Project", response.Message)

	})

	t.Run("ErrorDelete", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/projects/:id")

		projectController := New(&MockFalseProjectRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(projectController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}
		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})
}

type MockProjectRepository struct{}

func (m MockProjectRepository) Get() ([]entities.Project, error) {
	return []entities.Project{
		{Nama: "ProjectKu"},
	}, nil
}

func (m MockProjectRepository) GetById(project_id int) (entities.Project, error) {
	return entities.Project{Nama: "ProjectKu"}, nil
}

func (m MockProjectRepository) ProjectRegister(newProject entities.Project) (entities.Project, error) {
	return entities.Project{Nama: "ProjectKu"}, nil
}

func (m MockProjectRepository) Update(project_id int, newProject entities.Project) (entities.Project, error) {
	return entities.Project{Nama: "ProjectKu"}, nil
}

func (m MockProjectRepository) Delete(project_id int) error {
	return nil
}

type MockFalseProjectRepository struct{}

func (m MockFalseProjectRepository) Get() ([]entities.Project, error) {
	return nil, errors.New("False Project Object")
}
func (m MockFalseProjectRepository) GetById(project_id int) (entities.Project, error) {
	return entities.Project{}, errors.New("False Get Object")
}
func (m MockFalseProjectRepository) ProjectRegister(newProject entities.Project) (entities.Project, error) {
	return entities.Project{}, errors.New("False Register Object")
}
func (m MockFalseProjectRepository) Update(project_id int, newProject entities.Project) (entities.Project, error) {
	return entities.Project{}, errors.New("False Update Object")
}
func (m MockFalseProjectRepository) Delete(project_id int) error {
	return errors.New("False Delete Object")
}

//MOCK AUTH

type mockAuthRepository struct{}

func (ma mockAuthRepository) Login(email, Password string) (entities.User, error) {
	return entities.User{Email: "test@gmail.com", Password: "xyz"}, nil
}
