package model

type Usuarios struct {
	Usuario_id   int     `json:"usuario_id"`
	Created_at   string  `json:"created_at"`
	Updated_at   string  `json:"updated_at"`
	Delete_at    *string `json:"delete_at"`
	Nome         string  `json:"nome"`
	Email        string  `json:"email"`
	Status       string  `json:"status"`
	Senha        string  `json:"senha"`
	Uuid_usuario string  `json:"uuid"`
}
