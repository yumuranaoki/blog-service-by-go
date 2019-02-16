package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yumuranaoki/blog-service-by-go/model"
)

// DB is postgresql database
var DB *gorm.DB

// Init initialize database
func Init() {
	fmt.Printf("initialize database ...")
	var err error
	DB, err = gorm.Open("postgres", "host=db user=kneegorilla dbname=blog password=password sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	defer DB.Close()

	DB.LogMode(true)

	DB.AutoMigrate(&model.User{})
}
