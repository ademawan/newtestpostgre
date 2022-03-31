package routes

import (
	"newtestpostgre/configs"
	"newtestpostgre/delivery/controllers/auth"
	"newtestpostgre/delivery/controllers/project"
	"newtestpostgre/delivery/controllers/task"
	"newtestpostgre/delivery/controllers/user"
	"newtestpostgre/middlewares"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

func RegisterPath(e *echo.Echo, ac *auth.AuthController, uc *user.UserController, tc *task.TaskController, pc *project.ProjectController) {

	//=========================================================
	//ROUT USERS
	e.POST("users/login", ac.Login())
	e.POST("users/register", uc.UserRegister())
	eAuth := e.Group("")
	eAuth.Use(m.BasicAuth(middlewares.BusicAuth))
	eAuth.GET("users", uc.Get())
	eAuth.GET("users/:id", uc.GetById())
	eAuth.PUT("users/:id", uc.Update())
	eAuth.DELETE("users/:id", uc.Delete())

	//===========================================================
	//ROUTE TASK
	eTask := e.Group("todo/")
	eTask.Use(m.JWT([]byte(configs.JWT_SECRET)))
	eTask.POST("tasks/register", tc.TaskRegister())
	eTask.GET("tasks", tc.Get())
	eTask.GET("tasks/:id", tc.GetById())
	eTask.PUT("tasks/:id", tc.Update())
	eTask.DELETE("tasks/:id", tc.Delete())
	// e.POST("task/:id/completed", tc.TaskCompleted())
	// e.POST("task/:id/reopen", tc.TaskReopen())

	//===========================================================
	//ROUTE PROJECT
	eProject := e.Group("")
	e.POST("projects/register", pc.ProjectRegister())
	eProject.Use(m.JWT([]byte(configs.JWT_SECRET)))
	eProject.GET("projects", pc.Get())
	eProject.GET("projects/:id", pc.GetById())
	eProject.PUT("projects/:id", pc.Update())
	eProject.DELETE("projects/:id", pc.Delete())

}
