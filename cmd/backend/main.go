package main

import (
	"backend/internal/app"
	"log"

	_ "github.com/swaggo/echo-swagger"
)

// @title           Go Echo Template API
// @version         1.0
// @description     Go echo template API swagger documentation

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	application, err := app.NewApp()
	if err != nil {
		log.Fatalf("Application initialization failed: %v", err)
	}
	application.Run()
	application.RunServer()
}
