package controller

import (
	"fiber/golang_fiber/database"
	"fiber/golang_fiber/model"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func BlogList(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Blog list page",
	}
	db := database.DB
	var data []model.Blog

	db.Find(&data)

	context["data"] = data

	c.Status(200)
	return c.JSON(context)
}

func BlogCreate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Blog created",
	}
	data := new(model.Blog)

	if err := c.BodyParser(&data); err != nil {
		log.Println("Error in parsing request")
	}

	if err := database.DB.Create(&data).Error; err != nil {
		context["statusText"] = "error in blog uploading"
		return c.Status(fiber.StatusBadRequest).JSON(context)
	} else {
		context["data"] = data
	}

	return c.Status(http.StatusCreated).JSON(context)
}

func BlogUpdate(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Blog update",
	}

	id := c.Params("id")
	var blog model.Blog

	database.DB.First(&blog, id)

	if blog.ID == 0 {
		log.Println("blog not found")
		context["statusText"] = "blog not found"
		return c.Status(400).JSON(context)
	}

	if err := c.BodyParser(&blog); err != nil {
		log.Println("Error in parsing")

	}

	if updateErr := database.DB.Save(&blog).Error; updateErr != nil {
		log.Println("Error in saving data")
		context["statusText"] = "error in saving data"
	} else {
		context["updated_data"] = &blog
	}

	c.Status(200)
	return c.JSON(context)
}

func BlogDelete(c *fiber.Ctx) error {
	context := fiber.Map{
		"statusText": "ok",
		"message":    "Blog deleted succesfully",
	}
	id := c.Params("id")

	var blog model.Blog

	database.DB.First(&blog, id)

	if blog.ID == 0 {
		context["message"] = "Blog not found"
		return c.Status(400).JSON(context)
	}

	if err := database.DB.Delete(&blog).Error; err != nil {
		context["message"] = "something went wrong"
		return c.Status(500).JSON(context)
	}

	c.Status(200)

	return c.JSON(context)
}
