package utils

import (
	"fmt"
	"newtestpostgre/configs"
	"newtestpostgre/entities"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(config *configs.AppConfig) *gorm.DB {
	//postgres://nmqdckadrkgkvn:3431d43e0c453af757a272d2d3c940180f96fb115bcbc25a7e5f3ff8743fa090@ec2-3-229-161-70.compute-1.amazonaws.com:5432/ddb7g4irblf0rr

	DBURL := fmt.Sprintf("postgres://%s:%v@%s:%s/%s", config.Database.Username, config.Database.Password, config.Database.Address, config.Database.Port, config.Database.Name)
	db, err := gorm.Open(DBURL), &gorm.Config{})
	if err != nil {
		fmt.Printf("Cannot connect to database ")
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
	db.Migrator().DropTable(&entities.Task{})
	db.Migrator().DropTable(&entities.Project{})
	db.Migrator().DropTable(&entities.User{})
	db.AutoMigrate(&entities.User{})
	db.AutoMigrate(&entities.Project{})
	db.AutoMigrate(&entities.Task{})
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
