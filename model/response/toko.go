package response

type GetTokoResponse struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
	IDUser   int    `json:"user_id"`
}

type GetTokoByIDResponse struct {
	ID       int    `json:"id"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
}
