package model

type Participantes struct {
	Participante_id int     `json:"participante_id" `
	Usuario_id      int     `json:"usuario_id"`
	Created_at      string  `json:"created_at"`
	Updated_at      string  `json:"updated_at"`
	Delete_at       *string `json:"delete_at"`
	Compromisso_id  int     `json:"compromisso_id"`
	Usuario         []Usuarios
}
