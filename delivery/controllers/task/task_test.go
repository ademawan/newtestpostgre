package task

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

	t.Run("GetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		taskController := New(MockTaskRepository{})
		// taskController.Get()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}

		var response GetTasksResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 200, response.Code)

	})
	t.Run("ErrorGetUser", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/tasks")

		falsetaskController := New(MockFalseTaskRepository{})
		// falsetaskController.Get()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(falsetaskController.Get())(context); err != nil {
			log.Fatal(err)
			return
		}

		var response GetTaskResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)
	})

}

func TestGetById(t *testing.T) {
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
	t.Run("GetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockTaskRepository{})
		// taskController.GetById()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := GetTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Get Taks Id", response.Message)

	})
	t.Run("ErorGetById", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		falsetaskController := New(MockFalseTaskRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(falsetaskController.GetById())(context); err != nil {
			log.Fatal(err)
			return
		}

		var response GetTaskResponseFormat

		json.Unmarshal([]byte(res.Body.Bytes()), &response)
		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "not found", response.Message)
	})

}

func TestTaskRegister(t *testing.T) {
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

	t.Run("TaskRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Taskku",
			"priority":   1,
			"user_ID":    1,
			"project_ID": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/register")

		taskController := New(MockTaskRepository{})
		// taskController.TaskRegister()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.TaskRegister())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := RegisterTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, http.StatusCreated, response.Code)

	})
	t.Run("ErorTaskRegister", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"namagh":     "Taskku",
			"priority":   1,
			"user_ID":    1,
			"project_ID": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))

		res := httptest.NewRecorder()

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))

		context := e.NewContext(req, res)
		context.SetPath("/tasks/register")

		taskController := New(MockFalseTaskRepository{})
		// userController.TaskRegister()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.TaskRegister())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := RegisterTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 500, response.Code)
		assert.Equal(t, "There is some error on server", response.Message)

	})

	t.Run("TaskRegisterBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Taskku",
			"priority":   1,
			"user_ID":    1,
			"project_ID": "test",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/register")

		taskController := New(MockFalseTaskRepository{})
		// userController.TaskRegister()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.TaskRegister())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := RegisterTaskResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
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
		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})

	t.Run("Update Task", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":       "Taskku",
			"priority":   12,
			"user_id":    1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockTaskRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Update Task", response.Message)

	})

	t.Run("ErrorUpdateTask", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"name":       "Taskku",
			"priority":   12,
			"user_id":    1,
			"project_id": 1,
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockFalseTaskRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "not found", response.Message)

	})
	t.Run("UpdateBind", func(t *testing.T) {
		e := echo.New()
		requestBody, _ := json.Marshal(map[string]interface{}{
			"nama":     "Taskku",
			"priority": "12",
		})
		req := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer(requestBody))
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockFalseTaskRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.Update())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := UpdateResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 400, response.Code)
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
		context := e.NewContext(req, res)
		context.SetPath("users/login")

		authControl := auth.New(mockAuthRepository{})
		authControl.Login()(context)

		responses := auth.UserLoginResponseFormat{}
		json.Unmarshal([]byte(res.Body.Bytes()), &responses)

		jwtToken = responses.Token
		assert.Equal(t, responses.Message, "success login")
	})
	t.Run("DeleteTask", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockTaskRepository{})
		// taskController.Delete()(context)
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 200, response.Code)
		assert.Equal(t, "Success Delete Task", response.Message)

	})

	t.Run("ErrorDeleteTask", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		res := httptest.NewRecorder()
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", jwtToken))
		context := e.NewContext(req, res)
		context.SetPath("/tasks/:id")

		taskController := New(&MockFalseTaskRepository{})
		if err := m.JWT([]byte(configs.JWT_SECRET))(taskController.Delete())(context); err != nil {
			log.Fatal(err)
			return
		}

		response := DeleteResponseFormat{}

		json.Unmarshal([]byte(res.Body.Bytes()), &response)

		assert.Equal(t, 404, response.Code)
		assert.Equal(t, "not found", response.Message)

	})
}

type MockTaskRepository struct{}

func (m MockTaskRepository) Get() ([]entities.Task, error) {
	return []entities.Task{
		{Nama: "Taskku", Priority: 1, User_ID: 1, Project_ID: 1},
	}, nil
}

func (m MockTaskRepository) GetById(taskId int) (entities.Task, error) {
	return entities.Task{Nama: "Taskku", Priority: 1, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) TaskRegister(newTask entities.Task) (entities.Task, error) {
	return entities.Task{Nama: "Taskku", Priority: 1, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {
	return entities.Task{Nama: "Taskku", Priority: 1, User_ID: 1, Project_ID: 1}, nil
}

func (m MockTaskRepository) Delete(taskId int) error {
	return nil
}

type MockFalseTaskRepository struct{}

func (m MockFalseTaskRepository) Get() ([]entities.Task, error) {
	return nil, errors.New("False Task Object")
}
func (m MockFalseTaskRepository) GetById(taskId int) (entities.Task, error) {
	return entities.Task{}, errors.New("False Get Object")
}
func (m MockFalseTaskRepository) TaskRegister(newTask entities.Task) (entities.Task, error) {
	return entities.Task{}, errors.New("False Register Object")
}
func (m MockFalseTaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {
	return entities.Task{}, errors.New("False Update Object")
}
func (m MockFalseTaskRepository) Delete(taskId int) error {
	return errors.New("False Delete Object")
}

//MOCK AUTH

type mockAuthRepository struct{}

func (ma mockAuthRepository) Login(email, Password string) (entities.User, error) {
	return entities.User{Email: "test@gmail.com", Password: "xyz"}, nil
}
