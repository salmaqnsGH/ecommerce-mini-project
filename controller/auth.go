package controller

import (
	"fmt"
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/middleware"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"
	"mini-project-product/model/response"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(c *fiber.Ctx) error {
	user := new(request.RegisterUserRequest)
	if err := c.BodyParser(user); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	// TODO: validation duplicate email
	// TODO: validation duplicate phoneNumber

	var validate = validator.New()
	if err := validate.Struct(user); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	inputTime := user.TanggalLahir
	layoutIn := "02/01/2006"
	parsedTime, err := time.Parse(layoutIn, inputTime)
	if err != nil {
		return err
	}

	// TODO: hash password
	newUser := entity.User{
		Nama:         user.Nama,
		KataSandi:    user.KataSandi,
		NoTelp:       user.NoTelp,
		TanggalLahir: parsedTime,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
	}

	errCreateUser := db.DB.Create(&newUser).Error
	if errCreateUser != nil {
		var errors []string
		errors = append(errors, errCreateUser.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to POST data", nil, true, "Register Succeed")
	return c.JSON(response)
}

func LoginUser(c *fiber.Ctx) error {
	userRequest := new(request.LoginUserRequest)
	if err := c.BodyParser(userRequest); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	var validate = validator.New()
	if err := validate.Struct(userRequest); err != nil {
		errors := helper.FormatValidationError(err)
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	var user entity.User
	err := db.DB.First(&user, "no_telp = ?", userRequest.NoTelp).Error

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	if userRequest.KataSandi != user.KataSandi {
		var errors []string
		errors = append(errors, "password salah")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	fmt.Println("user", user)

	payload := entity.UserClaims{
		ID:      user.ID,
		Nama:    user.Nama,
		Email:   user.Email,
		NoTelp:  user.NoTelp,
		IsAdmin: user.IsAdmin,
	}
	token, err := middleware.EncodeJwt(payload)

	userResponse := response.LoginResponse{
		Nama:         user.Nama,
		NoTelp:       user.NoTelp,
		TanggalLahir: user.TanggalLahir,
		Tentang:      user.Tentang,
		Pekerjaan:    user.Pekerjaan,
		Email:        user.Email,
		IdProvinsi:   user.IdProvinsi,
		IdKota:       user.IdKota,
		Token:        token,
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, userResponse)
	return c.JSON(response)
}
