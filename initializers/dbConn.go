package initializers

import (
	"log"
	"os"

	"github.com/roshanlc/jwt-auth-go/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("SQLITE_DB")
	conn, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("error while connecting to db", err)
	}
	return conn
}

func InitialMigration() {
	db := ConnectDB()
	defer CloseDBConnection(db)
	db.AutoMigrate(models.User{})
}

func CloseDBConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
