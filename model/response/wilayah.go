package response

type GetListProvinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetListCitiesResponse struct {
	ID         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}
