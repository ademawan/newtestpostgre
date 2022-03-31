package task

import (
	"net/http"
	"newtestpostgre/delivery/controllers/auth"
	"newtestpostgre/delivery/controllers/common"
	"newtestpostgre/entities"
	"newtestpostgre/repository/task"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskController struct {
	repo task.Task
}

func New(repository task.Task) *TaskController {
	return &TaskController{
		repo: repository,
	}
}

func (tc *TaskController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := tc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Task", res))
	}
}

func (tc *TaskController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, _ := strconv.Atoi(c.Param("id"))

		res, err := tc.repo.GetById(taskId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Taks Id", res))
	}
}

func (tc *TaskController) TaskRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		task := RegisterTaskRequestFormat{}
		userId := int(auth.ExtractTokenUserId(c))
		if err := c.Bind(&task); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := tc.repo.TaskRegister(entities.Task{Nama: task.Nama, Priority: task.Priority, User_ID: userId, Project_ID: task.Project_ID, Status: -1})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Task", res))
	}
}

func (tc *TaskController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newTask = UpdateTaskRequestFormat{}
		taskId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newTask); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := tc.repo.Update(taskId, entities.Task{Nama: newTask.Nama, Priority: newTask.Priority})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Task", res))
	}
}

func (tc *TaskController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		taskId, _ := strconv.Atoi(c.Param("id"))

		err := tc.repo.Delete(taskId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Task", nil))
	}
}

// func (tc *TaskController) TaskCompleted() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		taskId, _ := strconv.Atoi(c.Param("id"))

// 		res, err := tc.repo.GetById(taskId)

// 		if err != nil {
// 			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
// 		}

// 		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Taks", res))
// 	}
// }

// func (tc *TaskController) TaskReopen() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		task := RegisterTaskRequestFormat{}

// 		if err := c.Bind(&task); err != nil {
// 			return c.JSON(http.StatusBadRequest, common.BadRequest())
// 		}

// 		res, err := tc.repo.TaskRegister(entities.Task{Nama: task.Nama, Priority: task.Priority, User_ID: task.User_ID, Project_ID: task.Project_ID})

// 		if err != nil {
// 			return c.JSON(http.StatusNotFound, common.InternalServerError())
// 		}

// 		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Task", res))
// 	}
// 	}
// }
