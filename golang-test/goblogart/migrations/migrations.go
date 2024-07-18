package main

import (
	"goblogart/inits"
	"goblogart/models"
)

func init() {
	inits.LoadEnv()
	inits.ConnectToDB()
}

func main() {
	inits.DB.AutoMigrate(&models.Post{})
	inits.DB.AutoMigrate(&models.User{})
}
