package controller

import (
	"log"
	"mini-project-product/db"
	"mini-project-product/helper"
	"mini-project-product/model/entity"
	"mini-project-product/model/request"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []entity.User
	result := db.DB.Find(&users)

	if result.Error != nil {
		log.Println(result.Error)
	}

	return c.JSON(users)
}

func RegisterUser(c *fiber.Ctx) error {
	user := new(request.RegisterUserRequest)
	if err := c.BodyParser(user); err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to POST data", errors, false, nil)

		return c.JSON(response)
	}

	// TODO: validation duplicate email

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
