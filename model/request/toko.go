package request

type UpdateTokoRequest struct {
	NamaToko string `form:"nama_toko"`
	URLFoto  string `form:"photo"`
}

type GetAllTokoRequest struct {
	NamaToko string `form:"nama_toko"`
	URLFoto  string `form:"photo"`
}
