package task

import (
	"newtestpostgre/entities"

	"gorm.io/gorm"
)

type TaskRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (tr *TaskRepository) Get() ([]entities.Task, error) {
	arrTask := []entities.Task{}

	if err := tr.database.Find(&arrTask).Error; err != nil {
		return nil, err
	}

	return arrTask, nil
}

func (tr *TaskRepository) GetById(taskId int) (entities.Task, error) {
	arrTask := entities.Task{}

	if err := tr.database.First(&arrTask, taskId).Error; err != nil {
		return arrTask, err
	}

	return arrTask, nil
}

func (tr *TaskRepository) TaskRegister(t entities.Task) (entities.Task, error) {
	if err := tr.database.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *TaskRepository) Update(taskId int, newTask entities.Task) (entities.Task, error) {

	var task entities.Task
	tr.database.First(&task, taskId)

	if err := tr.database.Model(&task).Updates(&newTask).Error; err != nil {
		return task, err
	}

	return task, nil
}

func (tr *TaskRepository) Delete(taskId int) error {

	var task entities.Task

	if err := tr.database.First(&task, taskId).Error; err != nil {
		return err
	}
	tr.database.Delete(&task, taskId)
	return nil

}
