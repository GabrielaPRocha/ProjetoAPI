package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type CompromissoRepository struct {
	connection *sql.DB
}

func NewCompromissoRepository(connection *sql.DB) *CompromissoRepository {
	return &CompromissoRepository{
		connection: connection,
	}
}

func (us *CompromissoRepository) GetCompromisso() ([]model.Compromisso, error) {
	query := "SELECT local_id,categoria_id,compromisso_id,horainicio,horafim FROM tb_compromisso"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Compromisso{}, err
	}
	var CompromissoList []model.Compromisso
	var CompromissoObj model.Compromisso

	for rows.Next() {
		err = rows.Scan(
			&CompromissoObj.Local_id,
			&CompromissoObj.Categoria_id,
			&CompromissoObj.Compromisso_id,
			/*			&CompromissoObj.Created_at,
						&CompromissoObj.Updated_at,
						&CompromissoObj.Delete_at,*/
			&CompromissoObj.Horainicio,
			&CompromissoObj.Horafim,
			//&CompromissoObj.Nome,
		//	&CompromissoObj.Participantes,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Compromisso{}, err
		}
		CompromissoList = append(CompromissoList, CompromissoObj)
	}
	rows.Close()
	return CompromissoList, nil
}
func (us *CompromissoRepository) GetCompromissoUsuario(usuarioUUID string) ([]model.Compromisso, error) {
	query := `SELECT local_id,categoria_id,p.compromisso_id,horainicio,horafim 
	FROM tb_compromisso c 
	JOIN tb_participantes p on c.compromisso_id = p.compromisso_id 
	JOIN tb_usuarios u on p.usuario_id = u.usuario_id
	WHERE u.uuid = ?`

	rows, err := us.connection.Query(query, usuarioUUID)
	if err != nil {
		fmt.Println(err)
		return []model.Compromisso{}, err
	}
	var CompromissoList []model.Compromisso
	var CompromissoObj model.Compromisso

	for rows.Next() {
		err = rows.Scan(
			&CompromissoObj.Local_id,
			&CompromissoObj.Categoria_id,
			&CompromissoObj.Compromisso_id,
			&CompromissoObj.Horainicio,
			&CompromissoObj.Horafim,
		)
		if err != nil {
			fmt.Println(err)
			return []model.Compromisso{}, err
		}
		fmt.Printf("%+v", CompromissoObj)
		CompromissoList = append(CompromissoList, CompromissoObj)
	}
	rows.Close()
	return CompromissoList, nil
}

func (us *CompromissoRepository) CreateCompromisso(compromisso model.Compromisso) (int, error) {
	query, err := us.connection.Prepare("INSERT INTO tb_compromisso (local_id,categoria_id,horainicio,horafim,participante) VALUES (?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(compromisso.Local_id, compromisso.Categoria_id, compromisso.Horainicio, compromisso.Horafim, compromisso.Participantes)
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
func (us *CompromissoRepository) DeleteCompromisso(compromisso model.Compromisso) (int, error) {
	query, err := us.connection.Prepare("UPDATE tb_compromisso SET deleted_at = ? WHERE compromisso_id = ?")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(compromisso.Delete_at, compromisso.Compromisso_id)
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
