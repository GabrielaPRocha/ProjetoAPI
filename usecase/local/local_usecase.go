package local

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

func (us *LocalUseCase) GetLocais() ([]model.Locais, error) {
	return us.repository.GetLocais()
}

func (us *LocalUseCase) CreateLocais(Local model.Locais) (model.Locais, error) {
	LocalId, err := us.repository.CreateLocais(Local)
	if err != nil {
		return model.Locais{}, err
	}
	Local.Local_id = LocalId
	return Local, nil
}
func (us *LocalUseCase) DeleteLocais(Local model.Locais) (model.Locais, error) {
	_, err := us.repository.DeleteLocais(Local)
	if err != nil {
		return model.Locais{}, err
	}
	return Local, nil
}
