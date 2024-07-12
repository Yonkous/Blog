package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Blog/handlers"
	"Blog/models"
	"Blog/router"
)

func main() {
	// Ortam değişkenlerinden veritabanı bağlantı bilgilerini alın
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	// Debugging: Ortam değişkenlerini yazdırın
	log.Println("DB_HOST:", dbHost)
	log.Println("DB_PORT:", dbPort)
	log.Println("DB_USER:", dbUser)
	log.Println("DB_NAME:", dbName)
	log.Println("DB_PASSWORD:", dbPassword)

	// DSN'i oluşturun
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPassword)

	// Debugging: DSN'i yazdırın
	log.Println("DSN:", dsn)

	// Veritabanı bağlantısını başlatın
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Basit bir sorgu çalıştırarak bağlantıyı test edin
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("failed to get generic database object: ", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal("failed to ping database: ", err)
	}

	// Veritabanı migrasyonlarını yapın
	models.Migrate(db)

	// Veritabanı bağlantısını handler'lara aktarın
	handlers.SetDB(db)

	// Router'ı oluşturun
	r := router.NewRouter()

	log.Println("Server started on :8082")
	log.Fatal(http.ListenAndServe(":8082", r))
}
