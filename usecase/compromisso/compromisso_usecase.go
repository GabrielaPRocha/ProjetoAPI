package compromisso

import (
	"fmt"
	"go-api/model"
	"go-api/repository"
	"go-api/usecase/participante"
)

type CompromissoUsecase struct {
	compromissorepository repository.CompromissoRepository
	participanteservice   participante.ParticipanteUsecase
	usuariorepository     repository.UsuarioRepository
}

func NewCompromissoUsecase(repo repository.CompromissoRepository, servic participante.ParticipanteUsecase, usuariorepo repository.UsuarioRepository) CompromissoUsecase {
	return CompromissoUsecase{
		compromissorepository: repo,
		participanteservice:   servic,
		usuariorepository:     usuariorepo,
	}
}

func (cu *CompromissoUsecase) GetCompromissos(usuarioUUID string) ([]model.Compromisso, error) {

	listadecompromissos, err := cu.compromissorepository.GetCompromissoUsuario(usuarioUUID)
	if err != nil {
		return []model.Compromisso{}, err
	}
	listadeparticipantes, err := cu.participanteservice.GetParticipantes()
	if err != nil {
		return []model.Compromisso{}, err
	}
	fmt.Printf("%+v\n,%+v\n", listadecompromissos, listadeparticipantes)
	return cu.Resolverparticipantes(listadecompromissos, listadeparticipantes)
}

func (cu *CompromissoUsecase) CreateCompromisso(compromisso model.Compromisso) (model.Compromisso, error) {
	compromissoID, err := cu.compromissorepository.CreateCompromisso(compromisso)
	if err != nil {
		return model.Compromisso{}, err
	}
	compromisso.Compromisso_id = compromissoID
	return compromisso, nil
}

func (cu *CompromissoUsecase) DeleteCompromisso(compromisso model.Compromisso) (model.Compromisso, error) {
	_, err := cu.compromissorepository.DeleteCompromisso(compromisso)
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
func (cu *CompromissoUsecase) UpdateParticipantesCompromisso(uuid_compromisso string, lista_emails_nova []string) ([]model.Compromisso, error) {
	// valida a lista de usuarios
	usuariosEncontrados, err := cu.usuariorepository.GetUsuariosByEmail(lista_emails_nova)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Emails encontrados: %v", usuariosEncontrados)
	//filtra emails validos

	if len(usuariosEncontrados) == 0 {
		return []model.Compromisso{}, fmt.Errorf("sem email valido")
	}
	var emailvalidado = []string{}
	var EmailInvalido = []string{}
	for _, email := range lista_emails_nova {
		emailValido := false
		for _, usuario := range usuariosEncontrados {
			if email == usuario.Email {
				emailValido = true
				emailvalidado = append(emailvalidado, email)
				break
			}
		}
		if !emailValido {
			EmailInvalido = append(EmailInvalido, email)
		}
	}
	fmt.Printf("\n\nemail inválido: %s", EmailInvalido)
	fmt.Printf("\nemail Valido: %s", emailvalidado)
	if len(EmailInvalido) > 0 {
		return nil, fmt.Errorf("email inválido: %s", EmailInvalido)
	}

	compromisso, err := cu.compromissorepository.GetCompromissoByUuid(uuid_compromisso)
	if err != nil {
		return nil, fmt.Errorf("Compromisso não encontrado: %v", err)
	}
	fmt.Printf("Compromisso encontrado: %v\n", compromisso)
	fmt.Print(compromisso)

	participantesAtuais, err := cu.participanteservice.GetParticipantesCompromisso(uuid_compromisso)
	if err != nil {
		return nil, fmt.Errorf("Erro ao obter participantes do compromisso: %v", err)
	}
	fmt.Printf("Participantes atuais: %v\n", participantesAtuais)
	participantesExistentes := make(map[string]bool)
	for _, participante := range participantesAtuais {
		participantesExistentes[participante.Usuario.Email] = true
	}
	// Remove todos os participantes atuais do compromisso
	for _, participante := range participantesAtuais {
		_, err := cu.participanteservice.DeleteParticipantes(participante)
		if err != nil {
			return nil, fmt.Errorf("Erro ao remover participante do compromisso: %v", err)
		}
	}

	// Adiciona novos participantes ao compromisso
	for _, email := range emailvalidado {
		usuarios, err := cu.usuariorepository.GetUsuariosByEmail([]string{email})
		if err != nil {
			return nil, fmt.Errorf("Erro ao buscar usuário por e-mail: %v", err)
		}
		if len(usuarios) == 0 {
			return nil, fmt.Errorf("Nenhum usuário encontrado com o e-mail: %s", email)
		}
		usuario := usuarios[0]

		novoParticipante := model.Participantes{
			Usuario_id:     usuario.Usuario_id,
			Compromisso_id: compromisso.Compromisso_id,
		}
		_, err = cu.participanteservice.CreateParticipante(novoParticipante)
		if err != nil {
			return nil, fmt.Errorf("Erro ao adicionar novo participante: %v", err)
		}
	}

	// Retorna o compromisso atualizado

	fmt.Printf("Compromisso atualizado: %v\n", compromisso)

	return []model.Compromisso{compromisso}, nil
}
