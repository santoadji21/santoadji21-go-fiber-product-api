package db

import (
	"fmt"
	"go-product-api/config"     // replace with your actual module path
	"go-product-api/pkg/models" // replace with your actual module path
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(cfg config.Config) {
    var err error

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        cfg.DBHost, cfg.DBUsername, cfg.DBPassword, cfg.DBName, cfg.DBPort)

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // AutoMigrate
    err = DB.AutoMigrate(&models.User{}, &models.Product{}, &models.Category{})
    if err != nil {
        log.Fatalf("Failed to auto-migrate: %v", err)
    }

     // Drop the unused column if it exists
    //  Db migrations users
    if DB.Migrator().HasColumn(&models.User{}, "OldColumn") {
        err = DB.Migrator().DropColumn(&models.User{}, "OldColumn")
        if err != nil {
            log.Fatalf("Failed to drop column: %v", err)
        }
    }

    // Db migrations products
    if DB.Migrator().HasColumn(&models.Product{}, "OldColumn") {
        err = DB.Migrator().DropColumn(&models.Product{}, "OldColumn")
        if err != nil {
            log.Fatalf("Failed to drop column: %v", err)
        }
    }

    // Db migrations categories
    if DB.Migrator().HasColumn(&models.Category{}, "OldColumn") {
        err = DB.Migrator().DropColumn(&models.Category{}, "OldColumn")
        if err != nil {
            log.Fatalf("Failed to drop column: %v", err)
        }
    }
}

func GetDB() *gorm.DB {
    return DB
}
