package model

type Locais struct {
	Local_id   int     `json:"local_id"`
	Created_at string  `json:"created_at"`
	Updated_at string  `json:"updated_at"`
	Delete_at  *string `json:"delete_at"`
	Nome       string  `json:"nome"`
	Presencial string  `json:"presencial"`
	Endereco   string  `json:"endereco"`
}
