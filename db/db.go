package db

import (
	"fmt"
	"os"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB // Deklarasi database

func InitDB() (*gorm.DB, error) {
	// Mengecek file .env
	if err := godotenv.Load(); 	err != nil {
		fmt.Println("Tidak ada file .env")
	}
	var err error

	dsn := os.Getenv("DATABASE_URL")

	DB, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Gagal terhubung ke database : ", err)
	}
	return DB, nil
}
