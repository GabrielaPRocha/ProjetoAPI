package model

type Categoria struct {
	Categoria_id int     `json:"categoria_id"`
	Created_at   string  `json:"created_at"`
	Updated_at   string  `json:"updated_at"`
	Delete_at    *string `json:"delete_at"`
	Nome         string  `json:"nome"`
	Slug         string  `json:"slug"`
}
