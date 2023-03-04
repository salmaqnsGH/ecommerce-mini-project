package controller

import (
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"
	"mini-project-product/model/response"

	"github.com/gofiber/fiber/v2"
)

func GetMyToko(c *fiber.Ctx) error {
	// TODO: fix idUser from jwt
	userId := 1
	var tokos []entity.Toko

	err := db.DB.Find(&tokos, "id_user = ?", userId).Error
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
	tokoID := c.Params("id")

	var toko entity.Toko

	err := db.DB.First(&toko, "id = ?", tokoID).Error

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
