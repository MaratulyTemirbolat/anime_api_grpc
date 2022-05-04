package db

import (
	"log"
	"proto/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbConfiguration string = "host = 127.0.0.1 port = 5432 user = postgres dbname = anime_api_trial password = 1538 sslmode=disable TimeZone=Asia/Almaty"

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(postgres.Open(dbConfiguration), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to db anime_api_trial: %v", err)
	}
	// Migration to database of created models
	database.AutoMigrate(
		&models.User{},
		&models.Phone{},
		&models.UserFriend{},
		&models.Action{},
		&models.Genre{},
		&models.Tag{},
		&models.Studio{},
		&models.AnimeGroup{},
		&models.Type{},

		&models.Anime{},
		&models.UserAnimeAction{},
		&models.Comment{},
	)
	DB = database
}
