package controller

import (
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"
	"mini-project-product/model/response"
	"time"

	"github.com/go-playground/validator/v10"
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

func UpdateProfile(c *fiber.Ctx) error {
	// TODO: fix idUser from jwt
	userRequest := new(request.UpdateProfileRequest)
	if err := c.BodyParser(userRequest); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	var validate = validator.New()
	if err := validate.Struct(userRequest); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	inputTime := userRequest.TanggalLahir
	layoutIn := "02/01/2006"
	parsedTime, err := time.Parse(layoutIn, inputTime)
	if err != nil {
		return err
	}

	// TODO: fix userId from wt
	var user entity.User
	userId := 3

	err = db.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	user.Nama = userRequest.Nama
	user.KataSandi = userRequest.KataSandi
	user.NoTelp = userRequest.NoTelp
	user.TanggalLahir = parsedTime
	user.Pekerjaan = userRequest.Pekerjaan
	user.Email = userRequest.Email
	user.IdProvinsi = userRequest.IdProvinsi
	user.IdKota = userRequest.IdKota

	errUpdate := db.DB.Save(&user).Error
	if errUpdate != nil {
		var errors []string
		errors = append(errors, errUpdate.Error())
		response := helper.APIResponse("Failed to UPDATE data", errors, false, nil)

		return c.JSON(response)
	}

	// TODO: match response to the requirements?
	userResponse := response.UpdateProfileResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: parsedTime,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		JenisKelamin: user.JenisKelamin,
		Tentang:      user.Tentang,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
	}

	response := helper.APIResponse("Succeed to UPDATE data", nil, true, userResponse)
	return c.JSON(response)
}

func GetAlamat(c *fiber.Ctx) error {
	// TODO: fix idUser from jwt
	userId := 1

	var alamats []entity.Alamat

	err := db.DB.Find(&alamats, "id_user = ?", userId).Error

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var alamatResponses []response.GetAlamatResponse

	for _, alamat := range alamats {
		alamatResponse := response.GetAlamatResponse{
			ID:           alamat.ID,
			JudulAlamat:  alamat.JudulAlamat,
			NoTelp:       alamat.NoTelp,
			NamaPenerima: alamat.NamaPenerima,
			DetailAlamat: alamat.DetailAlamat,
		}

		alamatResponses = append(alamatResponses, alamatResponse)
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, alamatResponses)
	return c.JSON(response)
}
