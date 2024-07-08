package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"

	"Blog/handlers"
	"Blog/models"
	"Blog/router"
)

func main() {
	// Veritabanı bağlantısını başlatın
	dsn := "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Veritabanı migrasyonlarını yapın
	models.Migrate(db)

	// Veritabanı bağlantısını handler'lara aktarın
	handlers.SetDB(db)

	// Router'ı oluşturun
	r := router.NewRouter()

	log.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
