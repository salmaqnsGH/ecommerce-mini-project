package response

type GetProvinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type GetCityResponse struct {
	ID         string `json:"id"`
	ProvinceId string `json:"province_id"`
	Name       string `json:"name"`
}
