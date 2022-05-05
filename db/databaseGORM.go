package db

import (
	"log"
	"proto/db/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dbConfigurat string = "host = 127.0.0.1 port = 5432 user = postgres dbname = trial_grpc password = 1538 sslmode=disable TimeZone=Asia/Almaty"

var DatabaseGORM *gorm.DB

func ConnectDBORM() {
	database, err := gorm.Open(postgres.Open(dbConfigurat), &gorm.Config{})
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
	err = database.SetupJoinTable(&models.User{}, "Friends", &models.UserFriend{})
	if err != nil {
		log.Fatalf("Извините, но у вас возникла ошибка: %v", err)
	}
	DatabaseGORM = database
}
