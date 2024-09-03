package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
	"fmt"
)

func main() {
	fmt.Println("Server running with Go")
	database.ConnectDB()
	routes.HandleRequest()
}
