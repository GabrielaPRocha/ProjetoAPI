package viewmodel

import (
	"go-api/model"
	"go-api/usecase/compromisso"
)

type CompromissoViewModel struct {
	CompromissoUsecase compromisso.CompromissoUsecase
}

func NewCompromissoViewModel(usecase compromisso.CompromissoUsecase) CompromissoViewModel {
	return CompromissoViewModel{
		CompromissoUsecase: usecase,
	}
}

func (vm *CompromissoViewModel) GetCompromisso() ([]model.Compromisso, error) {
	return vm.CompromissoUsecase.GetCompromissos()
}

func (vm *CompromissoViewModel) CreateCompromisso(compromissos model.Compromisso) (model.Compromisso, error) {
	return vm.CompromissoUsecase.CreateCompromisso(compromissos)
}

func (vm *CompromissoViewModel) DeleteCompromisso(compromissos model.Compromisso) (model.Compromisso, error) {
	return vm.CompromissoUsecase.DeleteCompromisso(compromissos)
}
