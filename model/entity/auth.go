package entity

type UserClaims struct {
	ID      int    `json:"id"`
	Nama    string `json:"nama"`
	Email   string `json:"email"`
	NoTelp  string `json:"noTelp"`
	IsAdmin bool   `json:"isAdmin"`
}
