package main

import (
	// "fmt"

	"github.com/orderzi/workout-service/server"
	// "github.com/orderzi/workout-service/utils"
)

// var fake_db []string

func main() {
	// name := utils.GetUserName()
	// age := utils.GetUserAge()
	// weight := utils.GetUserWeight()
	// user := utils.User{name, age, weight}
	// fmt.Println(user)
	// fmt.Println("Server started on port 8081")
	server.HandleRequests()

}
