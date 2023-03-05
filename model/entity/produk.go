package entity

import "time"

type Produk struct {
	ID            int       `json:"id"`
	IDToko        int       `json:"id_toko"`
	IDCategory    int       `json:"id_category"`
	NamaProduk    string    `json:"nama_produk"`
	Slug          string    `json:"slug"`
	HargaReseller string    `json:"harga_reseller"`
	HargaKonsumen string    `json:"harga_konsumen"`
	Stok          int       `json:"stok"`
	Deskripsi     string    `json:"deskripsi"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type FotoProduk struct {
	ID        int       `json:"id"`
	IDProduk  int       `json:"id_produk"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
