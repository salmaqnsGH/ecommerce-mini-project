package request

type CreateCategoryRequest struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}

type UpdateCategoryRequest struct {
	NamaCategory string `json:"nama_category" validate:"required"`
}
