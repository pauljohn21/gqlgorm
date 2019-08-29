package main

import (
	"gqlgorm/models"
	conn "gqlgorm/utils/database"
)
func main()  {
	db := conn.GetInstance()
	defer conn.Close()
	db.AutoMigrate(&models.User{},&models.Todo{},&models.Role{})
}