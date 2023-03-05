package controller

import (
	"fmt"
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/middleware"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"
	"mini-project-product/model/response"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetMyToko(c *fiber.Ctx) error {
	requestToken := c.Get("token")
	isValid, _, err := middleware.CheckValidToken(requestToken)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	if !isValid {
		var errors []string
		errors = append(errors, "Unauthorized")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	userData, err := middleware.GetUserData(c)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	userId := userData.ID
	var tokos []entity.Toko

	err = db.DB.Find(&tokos, "id_user = ?", userId).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var tokoResponses []response.GetTokoResponse

	for _, toko := range tokos {
		alamatResponse := response.GetTokoResponse{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto:  toko.URLFoto,
			IDUser:   toko.IDUser,
		}

		tokoResponses = append(tokoResponses, alamatResponse)
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, tokoResponses)
	return c.JSON(response)
}

func GetTokoByID(c *fiber.Ctx) error {
	requestToken := c.Get("token")
	isValid, _, err := middleware.CheckValidToken(requestToken)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	if !isValid {
		var errors []string
		errors = append(errors, "Unauthorized")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	tokoID := c.Params("id")

	var toko entity.Toko

	err = db.DB.First(&toko, "id = ?", tokoID).Error

	if err != nil {
		var errors []string
		errors = append(errors, "Toko tidak ditemukan")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	tokoResponse := response.GetTokoByIDResponse{
		ID:       toko.ID,
		NamaToko: toko.NamaToko,
		URLFoto:  toko.URLFoto,
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, tokoResponse)
	return c.JSON(response)
}

func UpdateToko(c *fiber.Ctx) error {
	requestToken := c.Get("token")
	isValid, _, err := middleware.CheckValidToken(requestToken)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	if !isValid {
		var errors []string
		errors = append(errors, "Unauthorized")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	userData, err := middleware.GetUserData(c)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	userId := userData.ID

	tokoRequest := new(request.UpdateTokoRequest)
	if err := c.BodyParser(tokoRequest); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	file, err := c.FormFile("photo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Missing image file")
	}

	var validate = validator.New()
	if err := validate.Struct(tokoRequest); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	var toko entity.Toko
	tokoID := c.Params("id")

	err = db.DB.First(&toko, "id = ?", tokoID).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	path := fmt.Sprintf("images/%s", file.Filename)
	err = c.SaveFile(file, path)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to save file", errors, false, nil)

		return c.JSON(response)
	}

	toko.IDUser = userId
	toko.NamaToko = tokoRequest.NamaToko
	toko.URLFoto = file.Filename
	errUpdate := db.DB.Save(&toko).Error
	if errUpdate != nil {
		var errors []string
		errors = append(errors, errUpdate.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to UPDATE data", nil, true, "Update toko succeed")
	return c.JSON(response)
}

func GetAllToko(c *fiber.Ctx) error {
	requestToken := c.Get("token")
	isValid, _, err := middleware.CheckValidToken(requestToken)
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	if !isValid {
		var errors []string
		errors = append(errors, "Unauthorized")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	queryLimit := c.Query("limit")
	queryPage := c.Query("page")
	queryNama := c.Query("nama")

	pageSize, err := strconv.Atoi(queryLimit)
	if err != nil {
		fmt.Println(err)
	}

	page, err := strconv.Atoi(queryPage)
	if err != nil {
		fmt.Println(err)
	}

	if pageSize < 0 {
		pageSize = 0
	}

	if page < 0 {
		page = 0
	}

	offset := (page - 1) * pageSize

	var tokos []entity.Toko

	err = db.DB.Where("nama_toko LIKE ?", "%"+queryNama+"%").Limit(pageSize).Offset(offset).Find(&tokos).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var tokoResponses []response.GetTokoResponse

	for _, toko := range tokos {
		alamatResponse := response.GetTokoResponse{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto:  toko.URLFoto,
			IDUser:   toko.IDUser,
		}

		tokoResponses = append(tokoResponses, alamatResponse)
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, tokoResponses)
	return c.JSON(response)
}
