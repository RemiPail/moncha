package db

import (
	"log"

	"github.com/RemiPail/moncha/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "postgres://rlrqhesiyvrscm:8fd3698ff8790fbc5079142727262d06683b0be2de8d18bf535f505da44718a5@ec2-52-212-228-71.eu-west-1.compute.amazonaws.com:5432/daf4vnlgjqdjj2"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.ShopItem{})

	return db
}
