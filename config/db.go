package config

import (
	"fmt"
	"user-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
  username := "root"
  password := "root"
  host := "tcp(127.0.0.1:3306)"
  database := "db_user"

  dsn := fmt.Sprintf("%v:%v@%v/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, database)

  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

  if err != nil {
    panic(err.Error())
  }

  db.AutoMigrate(&models.User{})

  return db
}