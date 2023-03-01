package entity

import "time"

type User struct {
	ID           int       `json:"id"`
	Nama         string    `json:"nama"`
	KataSandi    string    `json:"kata_sandi"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_lahir"`
	JenisKelamin string    `json:"jenis_kelamin"`
	Tentang      string    `json:"tentang"`
	Pekerjaan    string    `json:"pekerjaan"`
	Email        string    `json:"email"`
	IdProvinsi   string    `json:"id_provinsi"`
	IdKota       string    `json:"id_kota"`
	IsAdmin      bool      `json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
