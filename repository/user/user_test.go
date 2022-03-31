package user

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

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)

	t.Run("Success Create User", func(t *testing.T) {
		mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		res, err := repo.UserRegister(mockUser)
		assert.Equal(t, nil, err)
		assert.Equal(t, "Steven", res.Nama)
	})

	t.Run("Fail Create User", func(t *testing.T) {
		mockUser := entities.User{Model: gorm.Model{ID: 1}, Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		_, err := repo.UserRegister(mockUser)
		assert.NotNil(t, err)
	})
}

func TestGet(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Getting User", func(t *testing.T) {
		res, err := repo.Get()
		assert.Equal(t, nil, err)
		assert.Equal(t, mockUser.Nama, res[0].Nama)
	})

	db.Migrator().DropTable(&entities.User{})
	// db.AutoMigrate(&entities.User{})

	t.Run("Fail Getting User", func(t *testing.T) {
		_, err := repo.Get()
		assert.NotNil(t, err)
	})
}

func TestGetById(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Getting User ID", func(t *testing.T) {
		res, err := repo.GetById(int(mockUser.ID))
		assert.Equal(t, nil, err)
		assert.Equal(t, mockUser.Nama, res.Nama)
	})

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	t.Run("Fail Getting User ID", func(t *testing.T) {
		_, err := repo.GetById(int(mockUser.ID))
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Update User", func(t *testing.T) {
		mockUpdate := entities.User{Nama: "Testerr", Email: "steven@steven.com", Password: "steven123"}
		res, err := repo.Update(1, mockUpdate)
		assert.Equal(t, nil, err)
		assert.Equal(t, mockUpdate.Nama, res.Nama)
	})

	t.Run("Fail Update User", func(t *testing.T) {
		mockUpdate := entities.User{Model: gorm.Model{ID: 1}, Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
		_, err := repo.Update(2, mockUpdate)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Deleting User ID", func(t *testing.T) {
		err := repo.Delete(int(mockUser.ID))
		assert.Equal(t, nil, err)
	})

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	t.Run("Fail Getting User ID", func(t *testing.T) {
		err := repo.Delete(int(mockUser.ID))
		assert.NotNil(t, err)
	})
}
