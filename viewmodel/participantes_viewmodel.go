package viewmodel

import (
	"go-api/model"
	"go-api/usecase"
)

type ParticipantesViewModel struct {
	ParticipanteUseCase usecase.ParticipanteUsecase
}

func NewParticipantesViewModel(usecase usecase.ParticipanteUsecase) ParticipantesViewModel {
	return ParticipantesViewModel{
		ParticipanteUseCase: usecase,
	}
}

func (vm *ParticipantesViewModel) GetParticipantes() ([]model.Participantes, error) {
	return vm.ParticipanteUseCase.GetParticipantes()
}

func (vm *ParticipantesViewModel) CreateParticipante(participante model.Participantes) (model.Participantes, error) {
	return vm.ParticipanteUseCase.CreateParticipante(participante)
}

func (vm *ParticipantesViewModel) DeleteParticipante(participante model.Participantes) (model.Participantes, error) {
	return vm.ParticipanteUseCase.DeleteParticipantes(participante)
}
