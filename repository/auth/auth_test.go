package auth

import (
	"fmt"
	"newtestpostgre/configs"
	"newtestpostgre/entities"
	"newtestpostgre/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	config := configs.GetConfig()

	db := utils.InitDB(config)

	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})

	repo := New(db)
	mockUser := entities.User{Nama: "Steven", Email: "steven@steven.com", Password: "steven123"}
	db.Create(&mockUser)

	t.Run("Success Login", func(t *testing.T) {
		mockUser := entities.User{Email: "steven@steven.com", Password: "steven123"}
		res, err := repo.Login(mockUser.Email, mockUser.Password)
		fmt.Println(res)
		assert.Equal(t, nil, err)
		assert.Equal(t, "steven@steven.com", res.Email)
	})

	t.Run("Fail Login", func(t *testing.T) {

		mockUser := entities.User{Email: "steven@steven.com", Password: "steven13"}
		_, err := repo.Login(mockUser.Email, mockUser.Password)
		assert.NotNil(t, err)
	})

}
