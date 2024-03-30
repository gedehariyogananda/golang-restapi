package main

import (
	"fmt"
	db "test/golang/config"
	"github.com/gofiber/fiber/v2"
	routes "test/golang/routes"
)

func main() {
	fmt.Println("tes project")

	// connect databases 
	db.Connect()

	// create fiber app
	app := fiber.New()

	// import routes 
	routes.Route(app)
	
	// port
	app.Listen(":3000")


}
