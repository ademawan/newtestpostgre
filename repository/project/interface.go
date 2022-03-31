package project

import "newtestpostgre/entities"

type Project interface {
	Get() ([]entities.Project, error)
	GetById(projectId int) (entities.Project, error)
	ProjectRegister(newProject entities.Project) (entities.Project, error)
	Update(projectId int, newProject entities.Project) (entities.Project, error)
	Delete(projectId int) error
}
