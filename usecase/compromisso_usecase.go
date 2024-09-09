package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type CompromissoUsecase struct {
	repository repository.CompromissoRepository
}

func NewCompromissoUsecase(repo repository.CompromissoRepository) CompromissoUsecase {
	return CompromissoUsecase{
		repository: repo,
	}
}

func (cu *CompromissoUsecase) GetCompromissos() ([]model.Compromisso, error) {
	return cu.repository.GetCompromisso()
}

func (cu *CompromissoUsecase) CreateCompromisso(compromisso model.Compromisso) (model.Compromisso, error) {
	compromissoID, err := cu.repository.CreateCompromisso(compromisso)
	if err != nil {
		return model.Compromisso{}, err
	}
	compromisso.Compromisso_id = compromissoID
	return compromisso, nil
}

func (cu *CompromissoUsecase) DeleteCompromisso(compromisso model.Compromisso) (model.Compromisso, error) {
	_, err := cu.repository.DeleteCompromisso(compromisso)
	if err != nil {
		return model.Compromisso{}, err
	}
	return compromisso, nil
}
