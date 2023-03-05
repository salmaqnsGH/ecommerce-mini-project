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

func GetProdukByID(c *fiber.Ctx) error {
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

	produkID := c.Params("id")

	var produk entity.Produk

	err = db.DB.First(&produk, "id = ?", produkID).Error
	if err != nil {
		var errors []string
		errors = append(errors, "No Data Product")
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var toko entity.Toko
	err = db.DB.First(&toko, "id_user = ?", userId).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var fotoProduks []entity.FotoProduk
	err = db.DB.Find(&fotoProduks, "id_produk = ?", produkID).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var category entity.Category
	err = db.DB.First(&category, "id = ?", produk.IDCategory).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	var fotoProdukResponses []response.FotoProdukResponse
	for _, fotoProduk := range fotoProduks {
		fotoProdukResponse := response.FotoProdukResponse{
			IDProduk: fotoProduk.IDProduk,
			URL:      fotoProduk.URL,
		}
		fotoProdukResponses = append(fotoProdukResponses, fotoProdukResponse)
	}

	produkResponse := response.GetProdukResponse{
		ID:            produk.ID,
		NamaProduk:    produk.NamaProduk,
		Slug:          produk.Slug,
		HargaReseller: produk.HargaReseller,
		HargaKonsumen: produk.HargaKonsumen,
		Stok:          produk.Stok,
		Deskripsi:     produk.Deskripsi,
		Toko: response.GetTokoByIDResponse{
			ID:       toko.ID,
			NamaToko: toko.NamaToko,
			URLFoto:  toko.URLFoto,
		},
		Category: response.GetCategoryResponse{
			ID:           category.ID,
			NamaCategory: category.NamaCategory,
		},
		Photos: fotoProdukResponses,
	}

	response := helper.APIResponse("Succeed to GET data", nil, true, produkResponse)
	return c.JSON(response)
}

func GetAllProduct(c *fiber.Ctx) error {
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

	queryNamaProduk := c.Query("nama_produk")
	queryLimit := c.Query("limit")
	queryPage := c.Query("page")
	queryCategoryID := c.Query("category_id")
	queryTokoID := c.Query("toko_id")
	queryMaxHarga := c.Query("max_harga")
	queryMinHarga := c.Query("min_harga")

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

	var produks []entity.Produk

	err = db.DB.Where("nama_produk LIKE ? AND id_category = ? AND id_toko = ? AND harga_reseller <= ? AND harga_reseller <= ?", "%"+queryNamaProduk+"%", queryCategoryID, queryTokoID, queryMaxHarga, queryMinHarga).Limit(pageSize).Offset(offset).Find(&produks).Error
	if err != nil {
		var errors []string
		errors = append(errors, err.Error())
		response := helper.APIResponse("Failed to GET data", errors, false, nil)

		return c.JSON(response)
	}

	// var tokoResponses []response.GetTokoResponse

	// for _, toko := range tokos {
	// 	alamatResponse := response.GetTokoResponse{
	// 		ID:       toko.ID,
	// 		NamaToko: toko.NamaToko,
	// 		URLFoto:  toko.URLFoto,
	// 		IDUser:   toko.IDUser,
	// 	}

	// 	tokoResponses = append(tokoResponses, alamatResponse)
	// }

	response := helper.APIResponse("Succeed to GET data", nil, true, produks)
	return c.JSON(response)
}
