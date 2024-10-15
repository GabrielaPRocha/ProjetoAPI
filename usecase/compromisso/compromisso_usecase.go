package compromisso

import (
	"fmt"
	"go-api/model"
	"go-api/repository"
	"go-api/usecase/participante"
)

type CompromissoUsecase struct {
	repository repository.CompromissoRepository
	service    participante.ParticipanteUsecase
}

func NewCompromissoUsecase(repo repository.CompromissoRepository, servic participante.ParticipanteUsecase) CompromissoUsecase {
	return CompromissoUsecase{
		repository: repo,
		service:    servic,
	}
}

func (cu *CompromissoUsecase) GetCompromissos(usuarioUUID string) ([]model.Compromisso, error) {

	listadecompromissos, err := cu.repository.GetCompromissoUsuario(usuarioUUID)
	if err != nil {
		return []model.Compromisso{}, err
	}
	listadeparticipantes, err := cu.service.GetParticipantes()
	if err != nil {
		return []model.Compromisso{}, err
	}
	fmt.Printf("%+v\n,%+v\n", listadecompromissos, listadeparticipantes)
	return cu.resolverparticipantes(listadecompromissos, listadeparticipantes)
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
func (cu *CompromissoUsecase) resolverparticipantes(listacompromisso []model.Compromisso, listaparticipante []model.Participantes) ([]model.Compromisso, error) {
	mapcompromissos := make(map[int][]model.Participantes)
	for _, compromisso := range listacompromisso {
		//mapcompromissos[compromisso.Compromisso_id] =[]model.Participantes{}
		mapcompromissos[compromisso.Compromisso_id] = make([]model.Participantes, 0)
		for _, participante := range listaparticipante {
			if participante.Compromisso_id == compromisso.Compromisso_id {
				mapcompromissos[compromisso.Compromisso_id] = append(mapcompromissos[compromisso.Compromisso_id], participante)
			}
		}
	}
	compromissoscomparticipantes := listacompromisso
	//	compromissoscomparticipantes := listacompromisso
	for i, compromisso := range compromissoscomparticipantes {
		compromissoscomparticipantes[i].Participantes = mapcompromissos[compromisso.Compromisso_id]
	}
	fmt.Println(compromissoscomparticipantes)
	return compromissoscomparticipantes, nil
}
