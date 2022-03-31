package project

import (
	"newtestpostgre/entities"

	"gorm.io/gorm"
)

type ProjectRepository struct {
	database *gorm.DB
}

func New(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		database: db,
	}
}

func (tr *ProjectRepository) Get() ([]entities.Project, error) {
	arrProject := []entities.Project{}

	if err := tr.database.Find(&arrProject).Error; err != nil {
		return nil, err
	}

	return arrProject, nil
}

func (tr *ProjectRepository) GetById(projectId int) (entities.Project, error) {
	arrProject := entities.Project{}

	if err := tr.database.Preload("Task").Find(&arrProject, projectId).Error; err != nil {
		return arrProject, err
	}

	return arrProject, nil
}

func (tr *ProjectRepository) ProjectRegister(t entities.Project) (entities.Project, error) {
	if err := tr.database.Create(&t).Error; err != nil {
		return t, err
	}

	return t, nil
}

func (tr *ProjectRepository) Update(projectId int, newProject entities.Project) (entities.Project, error) {

	var project entities.Project
	tr.database.First(&project, projectId)

	if err := tr.database.Model(&project).Updates(&newProject).Error; err != nil {
		return project, err
	}

	return project, nil
}

func (tr *ProjectRepository) Delete(projectId int) error {

	var project entities.Project

	if err := tr.database.First(&project, projectId).Error; err != nil {
		return err
	}
	tr.database.Delete(&project, projectId)
	return nil

}
