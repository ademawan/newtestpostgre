package project

import (
	"newtestpostgre/entities"

	"gorm.io/gorm"
)

//----------------------------------------------------
//REQUEST FORMAT
//----------------------------------------------------
type RegisterProjectRequestFormat struct {
	Nama string `json:"nama" form:"nama"`
}
type CompletedProjectRequestFormat struct {
	Nama string `json:"nama" form:"nama"`
}
type ReopenProjectRequestFormat struct {
	Nama string `json:"nama" form:"nama"`
}
type UpdateProjectRequestFormat struct {
	gorm.Model
	Nama string `json:"nama" form:"nama"`
}

//-----------------------------------------------------
//RESPONSE FORMAT
//-----------------------------------------------------
type RegisterProjectResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Project `json:"data"`
}

type GetProjectsResponseFormat struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    []entities.Project `json:"data"`
}

type GetProjectResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Project `json:"data"`
}

type UpdateResponseFormat struct {
	Code    int              `json:"code"`
	Message string           `json:"message"`
	Data    entities.Project `json:"data"`
}

type DeleteResponseFormat struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
