package model

type Compromisso struct {
	Compromisso_id int
	Local_id       int
	Categoria_id   int
	Created_at     string
	Updated_at     string
	Delete_at      *string
	Horainicio     string
	Horafim        string
	Participantes  []Participantes
}
