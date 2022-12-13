package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/octane77/rova/identityService/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

func SetUpDatabaseConnection() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Sprintf("Error loading environment variables: %v", err))
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		// Keep trying to connect to DB
		time.Sleep(5 * time.Second)
		SetUpDatabaseConnection()
	}

	//TODO: RUN MIGRATIONS HERE
	db.Migrator().DropTable(entities.User{})
	db.Migrator().AutoMigrate(entities.User{})
	log.Println("Connected To PostgresDB!")
	return db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbInstance, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Error Closing Database: %v", err))
	}
	log.Println("Closing Database Connection")
	dbInstance.Close()
}
