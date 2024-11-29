package main

import (
	"log"

	_ "github.com/Jamescog/birth-buddy/docs"
	_ "github.com/mattn/go-sqlite3"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

//	@title			Birth Buddy API
//	@version		1.0
//	@description	API for Birth Buddy Birthday reminder application built with Golang and Ent ORM
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:8080
//	@BasePath	/api/v1

//	@securityDefinitions.basic	BasicAuth

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
func main() {

	server := NewServer()

	server.InitDatabase()

	r := server.http
	setupRoutes(r, server.db)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer server.CloseDatabase()

}
