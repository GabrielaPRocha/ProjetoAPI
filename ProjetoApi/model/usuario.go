package model

type Usuarios struct {
	usuario_id int    `json:"usuario_id"`
	created_at string `json:"created_at"`
	updated_at string `json:"updated_at"`
	delete_at  string `json:"delete_at"`
	nome       string `json:"nome"`
	email      string `json:"email"`
	status     bool   `json:"status"`
	senha      string `json:"senha"`
}
