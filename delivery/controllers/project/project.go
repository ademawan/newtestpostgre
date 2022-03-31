package project

import (
	"net/http"
	"newtestpostgre/delivery/controllers/common"
	"newtestpostgre/entities"
	"newtestpostgre/repository/project"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProjectController struct {
	repo project.Project
}

func New(repository project.Project) *ProjectController {
	return &ProjectController{
		repo: repository,
	}
}

func (tc *ProjectController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := tc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All Project", res))
	}
}

func (tc *ProjectController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		projectId, _ := strconv.Atoi(c.Param("id"))

		res, err := tc.repo.GetById(projectId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "Not Found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get Project", res))
	}
}

func (tc *ProjectController) ProjectRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		project := RegisterProjectRequestFormat{}

		if err := c.Bind(&project); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := tc.repo.ProjectRegister(entities.Project{Nama: project.Nama})

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create Project", res))
	}
}

func (tc *ProjectController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newProject = UpdateProjectRequestFormat{}
		projectId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newProject); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := tc.repo.Update(projectId, entities.Project{Nama: newProject.Nama})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update Project", res))
	}
}

func (tc *ProjectController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		projectId, _ := strconv.Atoi(c.Param("id"))

		err := tc.repo.Delete(projectId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete Project", nil))
	}
}
