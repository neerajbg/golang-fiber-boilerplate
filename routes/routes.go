package routes

import (
	"github.com/neerajbg/golang-fiber-boilerplate/controller"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", controller.HomePage)
	// app.Get("/blog", controller.BlogListPage)
	// app.Get("/blog/:id", controller.BlogPage)

}