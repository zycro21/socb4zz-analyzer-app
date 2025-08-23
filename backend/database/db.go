package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"social-analyzer-backend/models"
)

var DB *gorm.DB

func ConnectDB() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host, user, password, dbName, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db
	fmt.Println("✅ Database connected")
}

// Migrate all models
func Migrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.SocialAccount{},
		&models.Post{},
		&models.Comment{},
		&models.Engagement{},
		&models.SentimentAnalysis{},
		&models.Topic{},
		&models.PostTopic{},
		&models.Trend{},
		&models.Report{},
	)
	if err != nil {
		log.Fatal("❌ Migration failed:", err)
	}
	fmt.Println("✅ Database migrated")
}
