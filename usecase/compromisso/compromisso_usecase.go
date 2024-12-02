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
	return cu.Resolverparticipantes(listadecompromissos, listadeparticipantes)
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
func (cu *CompromissoUsecase) Resolverparticipantes(listacompromisso []model.Compromisso, listaparticipante []model.Participantes) ([]model.Compromisso, error) {
	//mapa para armazenar os compromissos e seus participantes
	mapcompromissos := make(map[int][]model.Participantes)
	// aqui inicio um loop que vai em cada elemnto de compromisso na lista
	for _, compromisso := range listacompromisso {
		//mapcompromissos[compromisso.Compromisso_id] =[]model.Participantes{}
		//inicia uma lista vazia, para cada compromisso cria uma lista vazia de participante no mapcompromisso
		mapcompromissos[compromisso.Compromisso_id] = make([]model.Participantes, 0)
		for _, participante := range listaparticipante {
			//verifica se pertence ao compromisso atual
			if participante.Compromisso_id == compromisso.Compromisso_id {
				//se sim, adiciona na lista de compromisso do mapa
				mapcompromissos[compromisso.Compromisso_id] = append(mapcompromissos[compromisso.Compromisso_id], participante)
			}
		}
	}
	//cria uma copia da lista original
	compromissoscomparticipantes := listacompromisso
	//	compromissoscomparticipantes := listacompromisso
	for i, compromisso := range compromissoscomparticipantes {
		compromissoscomparticipantes[i].Participantes = mapcompromissos[compromisso.Compromisso_id]
	}
	fmt.Println(compromissoscomparticipantes)
	return compromissoscomparticipantes, nil
}
func (cu *CompromissoUsecase) UpdateParticipanteCompromisso(uuidcompromisso string, lista_emails_nova []model.Participantes) ([]model.Compromisso, error) {
	// Obtém a lista atual de compromissos
	listacompromisso, err := cu.repository.GetCompromissoUsuario(uuidcompromisso)
	if err != nil {
		return []model.Compromisso{}, err
	} // Valida a lista de participantes e atualiza conforme necessário
	for i, compromisso := range listacompromisso {
		// Filtra os participantes que já estão no compromisso
		participantesExistentes := make(map[int]bool)
		for _, participante := range compromisso.Participantes {
			participantesExistentes[participante.Participante_id] = true
		} // Adiciona novos participantes se não estiverem já cadastrados
		for _, novoParticipante := range lista_emails_nova {
			if compromisso.Compromisso_id == novoParticipante.Compromisso_id {
				if !participantesExistentes[novoParticipante.Participante_id] {
					compromisso.Participantes = append(compromisso.Participantes, novoParticipante)
				}
			}
		} // Atualiza o compromisso na lista
		listacompromisso[i] = compromisso
	} // Atualiza os compromissos no repositório
	for _, compromisso := range listacompromisso {
		err := cu.repository.(compromisso)
		if err != nil {
			return []model.Compromisso{}, err
		}
	}
	return listacompromisso, nil
}
