package controller

import (
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"
	"mini-project-product/model/response"

	"github.com/gofiber/fiber/v2"
)

func GetProfile(c *fiber.Ctx) error {
	// TODO: fix idUser from jwt
	userId := 2

	var user entity.User

	err := db.DB.First(&user, "id = ?", userId).Error

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	userResponse := response.GetProfileResponse{
		ID:           user.ID,
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
		IsAdmin:      user.IsAdmin,
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, userResponse)
	return c.JSON(response)
}
