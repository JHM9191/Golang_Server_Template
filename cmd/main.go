package main

import (
	"Server/api/v1/router"
	_ "Server/cmd/docs"
	"Server/internal/util/db"
	"fmt"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 3.39.24.154:8000
// @BasePath /api/v1
func main() {
	database := db.NewDatabase()
	db.CreateAllTables(database)

	router.Server()

	ch := make(chan bool)
	fmt.Println("This is test")
	<-ch
}
