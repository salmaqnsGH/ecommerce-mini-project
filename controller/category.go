package controller

import (
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"

	"github.com/go-playground/validator/v10"
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

func CreateCategory(c *fiber.Ctx) error {
	// TODO: can only be accessed by admin
	category := new(request.CreateCategoryRequest)
	if err := c.BodyParser(category); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	var validate = validator.New()
	if err := validate.Struct(category); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	newCategory := entity.Category{
		NamaCategory: category.NamaCategory,
	}

	errCreateCategory := db.DB.Create(&newCategory).Error
	if errCreateCategory != nil {
		var errors []string
		errors = append(errors, errCreateCategory.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to POST data", nil, true, 1)
	return c.JSON(response)
}
