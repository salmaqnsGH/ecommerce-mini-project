package response

import "time"

type GetProfileResponse struct {
	ID           int       `json:"id"`
	Nama         string    `json:"nama"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string    `json:"email"`
	IdProvinsi   string    `json:"id_provinsi"`
	IdKota       string    `json:"id_kota"`
	IsAdmin      bool      `json:"is_admin"`
}

type UpdateProfileResponse struct {
	Nama         string    `json:"nama"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string    `json:"email"`
	IdProvinsi   string    `json:"id_provinsi"`
	IdKota       string    `json:"id_kota"`
	IsAdmin      bool      `json:"is_admin"`
}

type GetAlamatResponse struct {
	ID           int    `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}
