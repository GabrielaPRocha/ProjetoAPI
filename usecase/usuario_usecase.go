package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type UsuarioUsecase struct {
	repository repository.UsuarioRepository
}

func NewUsuariosUseCase(repo repository.UsuarioRepository) UsuarioUsecase {
	return UsuarioUsecase{
		repository: repo,
	}
}

func (su *UsuarioUsecase) GetUsuarios() ([]model.Usuarios, error) {
	return su.repository.GetUsuarios()
}
func (us *UsuarioUsecase) CreateUsuarios(usuario model.Usuarios) (model.Usuarios, error) {
	usuarioId, err := us.repository.CreateUsuarios(usuario)
	if err != nil {
		return model.Usuarios{}, err
	}
	usuario.Usuario_id = usuarioId
	return usuario, nil
}
