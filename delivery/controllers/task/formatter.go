package task

import (
	"newtestpostgre/entities"

	"gorm.io/gorm"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type RegisterTaskRequestFormat struct {
	Nama       string `json:"nama" form:"nama"`
	Priority   int    `json:"priority" form:"priority"`
	Project_ID int    `json:"project_id" form:"project_id"`
}

// type CompletedTaskRequestFormat struct {
// 	gorm.Model
// 	Nama       string `json:"nama" form:"nama"`
// 	Status     string `json:"status" form:"status"`
// }
type ReopenTaskRequestFormat struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Priority int    `json:"priority" form:"priority"`
	Status   int    `json:"status" form:"status"`
}
type UpdateTaskRequestFormat struct {
	gorm.Model
	Nama     string `json:"nama" form:"nama"`
	Priority int    `json:"priority" form:"priority"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type RegisterTaskResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type GetTasksResponseFormat struct {
	Code    int             `json:"code"`
	Message string          `json:"message"`
	Data    []entities.Task `json:"data"`
}

type GetTaskResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type UpdateResponseFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type DeleteResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ReopenTaskResponFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}

type CompleteTaskResponFormat struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    entities.Task `json:"data"`
}
