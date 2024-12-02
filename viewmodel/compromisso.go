package viewmodel

import (
	"regexp"
)

type InsertCompromisso struct {
	Local_id     int    `json:"local_id"`
	Categoria_id int    `json:"categoria_id"`
	Horainicio   string `json:"horainicio"`
	Horafim      string `json:"horafim"`
}

type GetCompromisso struct {
	Compromisso_id int     `json:"compromisso_id"`
	Local_id       int     `json:"local_id"`
	Categoria_id   int     `json:"categoria_id"`
	Created_at     string  `json:"created_at"`
	Updated_at     string  `json:"updated_at"`
	Delete_at      *string `json:"delete_at"`
	Horainicio     string  `json:"horainicio"`
	Horafim        string  `json:"horafim"`
	//Participantes []Participantes
}

type UpdateParticipantesCompromisso struct {
	Emails []string `json:"emails"`
}

func (vm UpdateParticipantesCompromisso) Validate() bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	emailPattern, _ := regexp.Compile(emailRegex)

	for _, email := range vm.Emails {
		if !emailPattern.MatchString(email) {
			// Retornar erro ao inves de bool
			return false
		}
	}

	return true
}
