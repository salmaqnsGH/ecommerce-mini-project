package request

type UpdateTokoRequest struct {
	NamaToko string `form:"nama_toko"`
	URLFoto  string `form:"photo"`
}

type GetAllTokoRequest struct {
	NamaToko string `form:"nama_toko"`
	URLFoto  string `form:"photo"`
}

type CreateTokoRequest struct {
	IDUser   int    `json:"id_user"`
	NamaToko string `json:"nama_toko"`
	URLFoto  string `json:"url_foto"`
}
