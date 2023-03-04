package controller

import (
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"
	"mini-project-product/model/response"

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

func GetCategoryById(c *fiber.Ctx) error {
	// TODO: can only be accessed by admin
	caetgoryId := c.Params("id")

	var category entity.Category

	err := db.DB.First(&category, "id = ?", caetgoryId).Error

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	categoryResponse := response.GetCategoryResponse{
		ID:           category.ID,
		NamaCategory: category.NamaCategory,
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, categoryResponse)
	return c.JSON(response)
}

func UpdateCategory(c *fiber.Ctx) error {
	// TODO: can only be accessed by admin
	categoryRequest := new(request.UpdateCategoryRequest)
	if err := c.BodyParser(categoryRequest); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	var validate = validator.New()
	if err := validate.Struct(categoryRequest); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	var category entity.Category
	categoryId := c.Params("id")

	err := db.DB.First(&category, "id = ?", categoryId).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	category.NamaCategory = categoryRequest.NamaCategory
	errUpdate := db.DB.Save(&category).Error
	if errUpdate != nil {
		var errors []string
		errors = append(errors, errUpdate.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to UPDATE data", nil, true, "")
	return c.JSON(response)
}

func DeleteCategory(c *fiber.Ctx) error {
	// TODO: can only be accessed by admin
	caetgoryId := c.Params("id")

	var category entity.Category

	err := db.DB.Where("id = ?", caetgoryId).First(&category).Delete(&category).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to DELETE data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to DELETE data", nil, true, "")
	return c.JSON(response)
}
