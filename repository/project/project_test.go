package project

import (
	"newtestpostgre/configs"
	"newtestpostgre/entities"
	"newtestpostgre/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestInsert(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{})
	db.AutoMigrate(&entities.Task{})

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := New(db)

	t.Run("Success Creating Project", func(t *testing.T) {
		mockProject := entities.Project{Nama: "Project1"}
		res, err := repo.ProjectRegister(mockProject)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Fail Creating Project", func(t *testing.T) {
		mockProject := entities.Project{Model: gorm.Model{ID: 1}, Nama: "Project1"}
		_, err := repo.ProjectRegister(mockProject)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := New(db)
	mockProject := entities.Project{Nama: "Project1"}
	db.Create(&mockProject)

	t.Run("Success Getting Project", func(t *testing.T) {
		res, err := repo.Get()
		assert.Equal(t, nil, err)
		assert.Equal(t, mockProject.Nama, res[0].Nama)
	})

	db.Migrator().DropTable(&entities.Project{})

	t.Run("Fail Getting Project", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := New(db)
	mockProject := entities.Project{Nama: "Project1"}
	db.Create(&mockProject)

	t.Run("Success Getting Project by ID", func(t *testing.T) {
		res, err := repo.GetById(int(mockProject.ID))
		assert.Equal(t, nil, err)
		assert.Equal(t, mockProject.Nama, res.Nama)
	})

	db.Migrator().DropTable(&entities.Project{})

	t.Run("Fail Getting Project by ID", func(t *testing.T) {
		_, err := repo.GetById(int(mockProject.ID))
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := New(db)
	mockProject := entities.Project{Nama: "Project1"}
	db.Create(&mockProject)

	t.Run("Success Update Project", func(t *testing.T) {
		mockUpdate := entities.Project{Nama: "Project1"}
		res, err := repo.Update(1, mockUpdate)
		assert.Equal(t, nil, err)
		assert.Equal(t, mockUpdate.Nama, res.Nama)
	})

	t.Run("Fail Update Project", func(t *testing.T) {
		mockUpdate := entities.Project{Model: gorm.Model{ID: 1}, Nama: "Project1"}
		_, err := repo.Update(2, mockUpdate)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Project{})
	db.AutoMigrate(&entities.Project{})

	repo := New(db)
	mockProject := entities.Project{Nama: "Project1"}
	db.Create(&mockProject)

	t.Run("Success Deleting Project ID", func(t *testing.T) {
		err := repo.Delete(int(mockProject.ID))
		assert.Equal(t, nil, err)
	})

	db.Migrator().DropTable(&entities.Project{})

	t.Run("Fail Deleting Project ID", func(t *testing.T) {
		err := repo.Delete(int(mockProject.ID))
		assert.NotNil(t, err)
	})
}
