package main

import (
	"fmt"

	"github.com/bigmuramura/simple-todo/app/models"
)

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/

	fmt.Println(models.Db)

	/*
		u := &models.User{}
		u.Name = "test"
		u.Email = "test@examlpe.com"
		u.PassWord = "testtest"
		fmt.Println(u)

		u.CreateUser()
	*/

	u, _ := models.GetUser(1)
	fmt.Println(u)
}
