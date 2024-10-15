package model

type Participantes struct {
	Participante_id int     `json:"participante_id" `
	Usuario_id      int     `json:"usuario_id"`
	Created_at      string  `json:"created_at"`
	Updated_at      string  `json:"updated_at"`
	Delete_at       *string `json:"delete_at"`
	Compromisso_id  int     `json:"compromisso_id"`
}

type Compromisso struct {
	Compromisso_id int     `json:"compromisso_id"`
	Local_id       int     `json:"local_id"`
	Categoria_id   int     `json:"categoria_id"`
	Created_at     string  `json:"created_at"`
	Updated_at     string  `json:"updated_at"`
	Delete_at      *string `json:"delete_at"`
	Horainicio     string  `json:"horainicio"`
	Horafim        string  `json:"horafim"`
	//Nome           string  `json:"nome"`
	Participantes []Participantes
}
