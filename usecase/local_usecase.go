package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type LocalUseCase struct {
	repository repository.LocalRepository
}

func NewLocalUseCase(repo *repository.LocalRepository) LocalUseCase {
	return LocalUseCase{
		repository: *repo,
	}
}

func (su *LocalUseCase) GetLocais() ([]model.Locais, error) {
	return su.repository.GetLocais()
}

func (us *LocalUseCase) CreateLocais(Local model.Locais) (model.Locais, error) {
	LocalId, err := us.repository.CreateLocais(Local)
	if err != nil {
		return model.Locais{}, err
	}
	Local.Local_id = LocalId
	return Local, nil
}
