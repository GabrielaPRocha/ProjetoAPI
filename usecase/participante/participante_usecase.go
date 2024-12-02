package participante

import (
	"go-api/model"
	"go-api/repository"
)

type ParticipanteUsecase struct {
	repository repository.ParticipanteRepository
}

func NewParticipanteUsecase(repo repository.ParticipanteRepository) ParticipanteUsecase {
	return ParticipanteUsecase{
		repository: repo,
	}
}

func (pu ParticipanteUsecase) GetParticipantes() ([]model.Participantes, error) {
	return pu.repository.GetParticipante()
}

func (pu *ParticipanteUsecase) CreateParticipante(participante model.Participantes) (model.Participantes, error) {
	participanteID, err := pu.repository.CreateParticipante(participante)
	if err != nil {
		return model.Participantes{}, err
	}
	participante.Participante_id = participanteID
	return participante, nil
}

func (pu *ParticipanteUsecase) DeleteParticipantes(participante model.Participantes) (model.Participantes, error) {
	_, err := pu.repository.DeleteParticipantes(participante)
	if err != nil {
		return model.Participantes{}, err
	}
	return participante, nil
}

func (pu *ParticipanteUsecase) UpdateParticipanteCompromisso(uuid string, newParticipantes []model.Participantes) (model.Compromisso, error) {
	compromisso, err := pu.repository.GetCompromissos(uuid)
	if err != nil {
		return model.Compromisso{}, err
	}
}
