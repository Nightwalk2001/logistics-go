package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"logistics-go/config"
	"logistics-go/handlers"
	"logistics-go/mongodb"
)

func main() {
	conf := config.Load()
	mongodb.Setup(conf.Uri)
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Content-Type",
	}))

	app.Post("/shipper", handlers.InsertShipper)

	app.Get("/shippers", handlers.GetShippers)

	app.Delete("/shipper/:id", handlers.DeleteShipper)

	app.Get("/arrears", handlers.GetArrears)

	log.Fatal(app.Listen(":3000"))

	// defer func() {
	// 	_ = client.Disconnect(ctx)
	// }()
}
