package route

import (
	"mini-project-product/controller"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(r *fiber.App) {
	r.Get("api/v1/user/all", controller.GetUsers)
	r.Post("api/v1/auth/register", controller.RegisterUser)
	r.Get("api/v1/category", controller.GetCategories)
	r.Get("api/v1/category/:id", controller.GetCategoryById)
	r.Post("api/v1/category", controller.CreateCategory)
	r.Put("api/v1/category/:id", controller.UpdateCategory)
	r.Delete("api/v1/category/:id", controller.DeleteCategory)
	r.Get("api/v1/provcity/listprovincies", controller.GetListProvince)
	r.Get("api/v1/provcity/listcities/:id", controller.GetListCities)
	r.Get("api/v1/provcity/detailprovince/:id", controller.GetDetailProvince)
	r.Get("api/v1/provcity/detailcity/:id", controller.GetDetailCity)
	r.Get("api/v1/user", controller.GetProfile)
	r.Put("api/v1/user", controller.UpdateProfile)

}
