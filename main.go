package main

import (
	"fmt"
	"todo_go/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// log.Println("test")
	fmt.Println(models.Db)

	u := &models.User{}
	u.Name = "test222"
	u.Email = "test@test.com"
	u.PassWord = "testest"
	fmt.Println(u)
	u.CreateUser()

}
