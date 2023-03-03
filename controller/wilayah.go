package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mini-project-product/helper"
	"mini-project-product/model/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetListProvince(c *fiber.Ctx) error {
	result, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	var provincies []response.GetListProvinceResponse
	err = json.Unmarshal(responseData, &provincies)

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, provincies)

	return c.JSON(response)
}

func GetListCities(c *fiber.Ctx) error {
	provinceId := c.Params("id")

	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceId)

	result, err := http.Get(url)

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	responseData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	var provincies []response.GetListCitiesResponse
	err = json.Unmarshal(responseData, &provincies)

	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	response := helper.APIResponse("Succeed to get data", nil, true, provincies)

	return c.JSON(response)
}
