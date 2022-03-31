package user

import (
	"net/http"
	"newtestpostgre/delivery/controllers/common"
	"newtestpostgre/entities"
	"newtestpostgre/repository/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	repo user.User
}

func New(repository user.User) *UserController {
	return &UserController{
		repo: repository,
	}
}

func (uc *UserController) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := uc.repo.Get()

		if err != nil {
			return c.JSON(http.StatusInternalServerError, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get All User", res))
	}
}

func (uc *UserController) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("id"))

		res, err := uc.repo.GetById(userId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.NotFound(http.StatusNotFound, "not found", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Get User", res))
	}
}

func (uc *UserController) UserRegister() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := RegisterRequestFormat{}

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.UserRegister(entities.User{Nama: user.Nama, Email: user.Email, Password: user.Password})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusCreated, common.Success(http.StatusCreated, "Success Create User", res))
	}
}

func (uc *UserController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser = UpdateRequestFormat{}
		userId, _ := strconv.Atoi(c.Param("id"))

		if err := c.Bind(&newUser); err != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest())
		}

		res, err := uc.repo.Update(userId, entities.User{Nama: newUser.Nama, Email: newUser.Email, Password: newUser.Password})

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Update User", res))
	}
}

func (uc *UserController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		userId, _ := strconv.Atoi(c.Param("id"))

		err := uc.repo.Delete(userId)

		if err != nil {
			return c.JSON(http.StatusNotFound, common.InternalServerError())
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "Success Delete User", nil))
	}
}
