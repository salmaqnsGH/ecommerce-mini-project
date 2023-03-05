package controller

import (
	"fmt"
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/middleware"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(c *fiber.Ctx) error {
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

	produkRequest := new(request.CreateProdukRequest)
	if err := c.BodyParser(produkRequest); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	file, err := c.FormFile("photos")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Missing image file")
	}

	var validate = validator.New()
	if err := validate.Struct(produkRequest); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	var toko entity.Toko
	err = db.DB.Where("id_user = ?", userId).First(&toko).Error
	if err != nil {
		var errors []string
		errors = append(errors, "Toko tidak ditemukan")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	produk := entity.Produk{
		IDToko:        toko.ID,
		IDCategory:    produkRequest.IDCategory,
		NamaProduk:    produkRequest.NamaProduk,
		Slug:          produkRequest.Slug,
		HargaReseller: produkRequest.HargaReseller,
		HargaKonsumen: produkRequest.HargaKonsumen,
		Stok:          produkRequest.Stok,
		Deskripsi:     produkRequest.Deskripsi,
	}

	errCreateProduk := db.DB.Create(&produk).Error
	if errCreateProduk != nil {
		var errors []string
		errors = append(errors, errCreateProduk.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

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

	var fotoProduk entity.FotoProduk
	fotoProduk.IDProduk = produk.ID
	fotoProduk.URL = file.Filename

	errCreateFotoProduk := db.DB.Create(&fotoProduk).Error
	if errCreateFotoProduk != nil {
		var errors []string
		errors = append(errors, errCreateFotoProduk.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to POST data", nil, true, 4)
	return c.JSON(response)
}
