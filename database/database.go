package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"rangguy/backend-api/config"
	"rangguy/backend-api/model"
)

var DB *gorm.DB

func InitDB() {
	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASS", "postgresql")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbName := config.GetEnv("DB_NAME", "db_santrikoding")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbName, dbPort)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database", err)
	}

	fmt.Println("Connected to database!")

	// migrasi
	err = DB.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("failed to migrate database", err)
	}

	fmt.Println("Migrated successfully!")
}
