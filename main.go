package main

import ( 
	"os"
	"github.com/neerajbg/golang-fiber-boilerplate/routes"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main(){
	port := os.Getenv("port")

	if port == ""{
		port = "3000"
	}

	app := fiber.New()

	routes.SetupRoutes(app)

	app.Listen(":"+port)
}
