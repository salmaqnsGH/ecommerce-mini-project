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

}
