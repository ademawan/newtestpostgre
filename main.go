package main

import (
	"newtestpostgre/configs"
	ac "newtestpostgre/delivery/controllers/auth"
	pc "newtestpostgre/delivery/controllers/project"
	"newtestpostgre/delivery/controllers/routes"
	tc "newtestpostgre/delivery/controllers/task"
	uc "newtestpostgre/delivery/controllers/user"
	authRepo "newtestpostgre/repository/auth"
	projectRepo "newtestpostgre/repository/project"
	taskRepo "newtestpostgre/repository/task"
	userRepo "newtestpostgre/repository/user"
	"newtestpostgre/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	authRepo := authRepo.New(db)
	userRepo := userRepo.New(db)
	taskRepo := taskRepo.New(db)
	projectRepo := projectRepo.New(db)

	authController := ac.New(authRepo)
	userController := uc.New(userRepo)
	taskController := tc.New(taskRepo)
	projectController := pc.New(projectRepo)

	e := echo.New()

	routes.RegisterPath(e, authController, userController, taskController, projectController)

	log.Fatal(e.Start(":8080"))
}
