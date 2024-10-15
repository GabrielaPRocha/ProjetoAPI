package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ParticipanteRepository struct {
	connection *sql.DB
}

func NewParticipanteRepository(connection *sql.DB) *ParticipanteRepository {
	return &ParticipanteRepository{
		connection: connection,
	}
}

func (us *ParticipanteRepository) GetParticipante() ([]model.Participantes, error) {
	query := `SELECT idparticipantes, created_at, updated_at, delete_at, usuario_id,compromisso_id FROM tb_participantes `
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Participantes{}, err
	}
	var ParticipantesList []model.Participantes
	var ParticipantesObj model.Participantes

	for rows.Next() {
		err = rows.Scan(
			&ParticipantesObj.Participante_id,
			&ParticipantesObj.Created_at,
			&ParticipantesObj.Updated_at,
			&ParticipantesObj.Delete_at,
			&ParticipantesObj.Usuario_id,
			&ParticipantesObj.Compromisso_id,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Participantes{}, err
		}
		ParticipantesList = append(ParticipantesList, ParticipantesObj)
	}
	rows.Close()
	return ParticipantesList, nil
}

func (us *ParticipanteRepository) CreateParticipante(participante model.Participantes) (int, error) {
	query, err := us.connection.Prepare("INSERT INTO tb_participante (usuario_id, compromisso_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(participante.Usuario_id, participante.Compromisso_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(id), nil
}
func (us *ParticipanteRepository) DeleteParticipantes(participante model.Participantes) (int, error) {
	query, err := us.connection.Prepare("UPDATE tb_participante SET deleted_at = ? WHERE participante_id = ?")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(participante.Delete_at, participante.Participante_id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return int(rowsAffected), nil
}
