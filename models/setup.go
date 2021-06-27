package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=database-1.cngy27fgvclx.us-east-2.rds.amazonaws.com user=postgres password=trewq123 dbname=gin_crud port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("====>>> I think it works")
	}

	database.AutoMigrate(&Book{})

	DB = database

}
