package main

import (
	"user-api/config"
	"user-api/routes"
)
func main() {
  db := config.ConnectDataBase()
  sqlDB, _ := db.DB()
  defer sqlDB.Close()

  r := routes.SetupRouter(db)
  r.Run()
}