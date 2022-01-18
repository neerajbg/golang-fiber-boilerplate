package controller

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/golang-fiber-boilerplate/model"
	"github.com/neerajbg/golang-fiber-boilerplate/database"
)

func HomePage(c *fiber.Ctx) error{
	db := database.DB

	var blog []model.Blog
	db.Find(&blog)

	context := fiber.Map{
		"blog": blog,
	}

	c.Status(200)
	return c.JSON(context)
}

func CreateBlogPage(c *fiber.Ctx) error {
	blog := new(model.Blog)

	if err := c.BodyParser(blog); err != nil{
		log.Println("Error in parsing")
		return c.Status(500).JSON(fiber.Map{
			"success": false,
		})
	}

	// Now add post and return 
	database.DB.Create(blog)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"post": blog,
	})

}