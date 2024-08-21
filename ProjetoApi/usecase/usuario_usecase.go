package usecase

import (
	"ProjetoAPI/model"
)

type UsuarioUsecase struct {
}

func NewUsuariosUseCase() UsuarioUsecase {
	return UsuarioUsecase{}
}

func (us *UsuarioUsecase) GetUsuarios() ([]model.Usuario, error) {
	return []model.Usuario{}, nil
}
