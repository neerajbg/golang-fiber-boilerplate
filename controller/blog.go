package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/neerajbg/golang-fiber-boilerplate/model"
	"github.com/neerajbg/golang-fiber-boilerplate/database"
)
var db = database.Connect()

func HomePage(c *fiber.Ctx) error{

	var blog []model.Blog

	db.Find(&blog)

	context := fiber.Map{
		"blog": blog,
	}

	return c.Status(200).JSON(context)
}