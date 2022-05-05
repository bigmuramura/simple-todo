package main

import (
	"fmt"

	"github.com/bigmuramura/simple-todo/app/controllers"
	"github.com/bigmuramura/simple-todo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()
}
