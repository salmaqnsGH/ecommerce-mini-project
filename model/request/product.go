package request

type CreateProdukRequest struct {
	IDCategory    int    `form:"category_id" validate:"required"`
	NamaProduk    string `form:"nama_produk" validate:"required"`
	Slug          string `form:"slug"`
	HargaReseller string `form:"harga_reseller" validate:"required"`
	HargaKonsumen string `form:"harga_konsumen" validate:"required"`
	Stok          int    `form:"stok" validate:"required"`
	Deskripsi     string `form:"deskripsi" validate:"required"`
}

type FotoProdukRequest struct {
	IDProduk int    `json:"id_produk"`
	URL      string `json:"url"`
}
