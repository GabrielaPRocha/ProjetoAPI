package viewmodel

import (
	"go-api/model"
	"go-api/usecase/participante"
)

type ParticipantesViewModel struct {
	ParticipanteUseCase participante.ParticipanteUsecase
}

func NewParticipantesViewModel(usecase participante.ParticipanteUsecase) ParticipantesViewModel {
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
