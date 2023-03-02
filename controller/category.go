package controller

import (
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	// TODO: can only be accessed by admin
	var categories []entity.Category
	result := db.DB.Find(&categories)

	if result.Error != nil {
		var errors []string
		errors = append(errors, result.Error.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, categories)
	return c.JSON(response)
}
