package response

type GetProdukResponse struct {
	ID            int    `form:"id"`
	NamaProduk    string `form:"nama_produk"`
	Slug          string `form:"slug"`
	HargaReseller string `form:"harga_reseller"`
	HargaKonsumen string `form:"harga_konsumen"`
	Stok          int    `form:"stok"`
	Deskripsi     string `form:"deskripsi"`
	Toko          GetTokoByIDResponse
	Category      GetCategoryResponse
	Photos        []FotoProdukResponse
}

type FotoProdukResponse struct {
	IDProduk int    `json:"id_produk"`
	URL      string `json:"url"`
}
