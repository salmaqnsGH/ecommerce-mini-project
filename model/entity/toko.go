package entity

import "time"

type Toko struct {
	ID        int       `json:"id"`
	IDUser    int       `json:"id_user"`
	NamaToko  string    `json:"nama_toko"`
	URLFoto   string    `json:"url_foto"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
