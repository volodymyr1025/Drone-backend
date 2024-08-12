package database

import (
	"drone-backend/internal/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() error {

    env_err := godotenv.Load()
	if env_err != nil {
		log.Fatalf("Error loading .env file: %v", env_err)
	}

    host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

    // if host == "" {
	// 	host = "localhost"
	// }
	// if user == "" {
	// 	user = "postgres"
	// }
	// if password == "" {
	// 	password = "123456"
	// }
	// if dbname == "" {
	// 	dbname = "dronedb"
	// }
	// if port == "" {
	// 	port = "5432"
	// }
	// if sslmode == "" {
	// 	sslmode = "disable"
	// }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
    
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    // Drop and recreate the tables to ensure proper schema
    // err = DB.Migrator().DropTable(&models.Drone{}, &models.Attribute{}, &models.Policy{})
    // if err != nil {
    //     return err
    // }

    err = DB.AutoMigrate(
        &models.Drone{},
        &models.Attribute{},
        &models.Policy{},
        &models.History{},
    )
    if err != nil {
        return err
    }

    return nil
}
