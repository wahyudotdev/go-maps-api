package main

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	_ "go-maps-api/docs"
	"go-maps-api/router"
	"log"
	"os"
)

// @title Maps API
// @version 1.0
// @description Public API
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
func main() {
	app := fiber.New()
	app.Get("/docs/*", swagger.HandlerDefault)
	group := app.Group("/api")
	router.MapsRouter(group)
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
