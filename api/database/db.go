package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DBC = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", "localhost", "5432", "postgres", "api_go", "desarrollo")

func Connect() *gorm.DB {
	db, err := gorm.Open("postgres", DBC)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	// defer db.Close()

	return db
}