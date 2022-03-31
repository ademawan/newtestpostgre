package utils

import (
	"fmt"
	"newtestpostgre/configs"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {

	DBURL := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta ", config.Database.Address, config.Database.Username, config.Database.Password, config.Database.Port, config.Database.Name)
	db, err := gorm.Open(postgres.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to database")
	} else {
		fmt.Printf("We are connected to the %s database", DBURL)
	}

	if err != nil {
		log.Info("failed to connect database :", err)
		panic(err)
	}
	//fdkdkd

	InitMigrate(db)
	return db

}

func InitMigrate(db *gorm.DB) {
	// db.Migrator().DropTable(&entities.Task{})
	// db.Migrator().DropTable(&entities.Project{})
	// // db.Migrator().DropTable(&entities.User{})
	// // db.AutoMigrate(&entities.User{})
	// db.AutoMigrate(&entities.Project{})
	// db.AutoMigrate(&entities.Task{})
	// // for i := 0; i < 2000; i++ {
	// // 	db.Create(&entities.User{
	// // 		Nama:     faker.Name(),
	// // 		Email:    faker.Email(),
	// // 		Password: "xyz",
	// // 	})
	// // }
	// for i := 0; i < 500; i++ {
	// 	db.Create(&entities.Project{
	// 		Nama: faker.TitleMale(),
	// 	})
	// }
	// for i := 0; i < 500; i++ {
	// 	db.Create(&entities.Task{
	// 		Nama:       faker.TitleMale(),
	// 		User_ID:    int(math.Round(float64(rand.Intn(20)))),
	// 		Project_ID: int(math.Round(float64(rand.Intn(100)))),
	// 	})
	// }
}
