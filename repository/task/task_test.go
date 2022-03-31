package task

import (
	"newtestpostgre/configs"
	"newtestpostgre/entities"
	projectRepo "newtestpostgre/repository/project"
	userRepo "newtestpostgre/repository/user"
	"newtestpostgre/utils"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// func Dumy() {
// 	config := configs.GetConfig()

// 	db := utils.InitDB(config)

// 	db.Migrator().DropTable(&entities.User{}, &entities.Project{})
// 	db.AutoMigrate(&entities.User{}, &entities.Project{})
// 	userRepo := userRepo.New(db)
// 	projectRepo := projectRepo.New(db)

// 	mockUser := entities.User{Nama: "Steven", Email: "test@gmail.com", Password: "test"}
// 	userRepo.UserRegister(mockUser)
// 	mockProject := entities.Project{Nama: "Steven"}
// 	projectRepo.ProjectRegister(mockProject)
// }
func TestInsert(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{}, &entities.User{}, &entities.Project{})
	db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.Project{})
	userRepo := userRepo.New(db)
	projectRepo := projectRepo.New(db)
	repo := New(db)

	mockUser := entities.User{Nama: "Steven", Email: "test@gmail.com", Password: "test"}
	userRepo.UserRegister(mockUser)
	mockProject := entities.Project{Nama: "Steven"}
	projectRepo.ProjectRegister(mockProject)

	t.Run("Success Creating Task", func(t *testing.T) {
		mockTask := entities.Task{Nama: "Steven", Priority: 1, User_ID: 1, Project_ID: 1, Status: -1}
		res, err := repo.TaskRegister(mockTask)
		assert.Equal(t, nil, err)
		assert.Equal(t, 1, int(res.ID))
	})

	t.Run("Fail Creating Task", func(t *testing.T) {
		mockTask := entities.Task{Model: gorm.Model{ID: 1}, Nama: "Steven", Priority: 1, User_ID: 1, Project_ID: 1, Status: -1}
		_, err := repo.TaskRegister(mockTask)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)
	db.Migrator().DropTable(&entities.Task{}, &entities.User{}, &entities.Project{})
	db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.Project{})
	userRepo := userRepo.New(db)
	projectRepo := projectRepo.New(db)
	repo := New(db)

	mockUser := entities.User{Nama: "Steven", Email: "test@gmail.com", Password: "test"}
	userRepo.UserRegister(mockUser)
	mockProject := entities.Project{Nama: "Steven"}
	projectRepo.ProjectRegister(mockProject)
	mockTask := entities.Task{Nama: "TestTask", Priority: 1, User_ID: 1, Project_ID: 1, Status: -1}
	db.Create(&mockTask)

	t.Run("Success Getting Task", func(t *testing.T) {
		res, err := repo.Get()
		assert.Equal(t, nil, err)
		assert.Equal(t, "TestTask", res[0].Nama)
	})

	db.Migrator().DropTable(&entities.Task{})

	t.Run("Fail Getting Task", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{}, &entities.User{}, &entities.Project{})
	db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.Project{})
	userRepo := userRepo.New(db)
	projectRepo := projectRepo.New(db)
	repo := New(db)

	mockUser := entities.User{Nama: "Steven", Email: "test@gmail.com", Password: "test"}
	userRepo.UserRegister(mockUser)
	mockProject := entities.Project{Nama: "Steven"}
	projectRepo.ProjectRegister(mockProject)

	mockTask := entities.Task{Nama: "Mawan", Priority: 1, User_ID: 1, Project_ID: 1, Status: -1}
	db.Create(&mockTask)
	t.Run("Success Getting Task by ID", func(t *testing.T) {
		res, err := repo.GetById(int(mockTask.ID))
		assert.Equal(t, nil, err)
		assert.Equal(t, mockTask.Nama, res.Nama)
	})

	db.Migrator().DropTable(&entities.Task{})

	t.Run("Fail Getting Task by ID", func(t *testing.T) {
		_, err := repo.GetById(int(mockTask.ID))
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{}, &entities.User{}, &entities.Project{})
	db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.Project{})
	userRepo := userRepo.New(db)
	projectRepo := projectRepo.New(db)
	repo := New(db)

	mockUser := entities.User{Nama: "Steven", Email: "test@gmail.com", Password: "test"}
	userRepo.UserRegister(mockUser)
	mockProject := entities.Project{Nama: "Steven"}
	projectRepo.ProjectRegister(mockProject)
	mockTask := entities.Task{Nama: "Task1", Priority: 12, User_ID: 1, Project_ID: 1, Status: -1}
	db.Create(&mockTask)

	t.Run("Success Update Task", func(t *testing.T) {
		mockUpdate := entities.Task{Nama: "Task1", Priority: 12, User_ID: 1, Project_ID: 1, Status: -1}
		res, err := repo.Update(1, mockUpdate)
		assert.Equal(t, nil, err)
		assert.Equal(t, mockUpdate.Nama, res.Nama)
	})

	t.Run("Fail Update Task", func(t *testing.T) {
		mockUpdate := entities.Task{Model: gorm.Model{ID: 1}, Nama: "Task1", Priority: 12, User_ID: 1, Project_ID: 1, Status: -1}
		_, err := repo.Update(2, mockUpdate)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.Task{}, &entities.User{}, &entities.Project{})
	db.AutoMigrate(&entities.Task{}, &entities.User{}, &entities.Project{})
	userRepo := userRepo.New(db)
	projectRepo := projectRepo.New(db)
	repo := New(db)

	mockUser := entities.User{Nama: "Steven", Email: "test@gmail.com", Password: "test"}
	userRepo.UserRegister(mockUser)
	mockProject := entities.Project{Nama: "Steven"}
	projectRepo.ProjectRegister(mockProject)
	mockTask := entities.Task{Nama: "Task1", Priority: 12, User_ID: 1, Project_ID: 1, Status: -1}
	db.Create(&mockTask)

	t.Run("Success Deleting Task ID", func(t *testing.T) {
		err := repo.Delete(int(mockTask.ID))
		assert.Equal(t, nil, err)
	})

	db.Migrator().DropTable(&entities.Task{})

	t.Run("Fail Deleting Task ID", func(t *testing.T) {
		err := repo.Delete(int(mockTask.ID))
		assert.NotNil(t, err)
	})
}
