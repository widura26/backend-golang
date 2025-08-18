package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Sesuaikan dengan user, password, host, port, dan nama database kamu
	dsn := "root:Widura260503@tcp(127.0.0.1:3307)/go_rest_api?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal connect ke database: ", err)
	}

	fmt.Println("✅ Berhasil connect ke database")
}
