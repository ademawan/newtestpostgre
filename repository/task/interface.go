package task

import "newtestpostgre/entities"

type Task interface {
	Get() ([]entities.Task, error)
	GetById(taskId int) (entities.Task, error)
	TaskRegister(newTask entities.Task) (entities.Task, error)
	Update(taskId int, newTask entities.Task) (entities.Task, error)
	Delete(taskId int) error
}
