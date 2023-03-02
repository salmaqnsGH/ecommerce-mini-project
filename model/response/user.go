package response

type GetCategoryByIdResponse struct {
	ID           int    `json:"id"`
	NamaCategory string `json:"nama_category"`
}
