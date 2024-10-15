package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type LocalRepository struct {
	connection *sql.DB
}

func NewLocalRepository(connection *sql.DB) *LocalRepository {
	return &LocalRepository{
		connection: connection,
	}
}

func (us *LocalRepository) GetLocais() ([]model.Locais, error) {
	query := "SELECT local_id, created_at, updated_at, delete_at, nome, presencial, endereco FROM tb_local"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Locais{}, err
	}
	var LocaisList []model.Locais
	var LocaisObj model.Locais

	for rows.Next() {
		err = rows.Scan(
			&LocaisObj.Local_id,
			&LocaisObj.Created_at,
			&LocaisObj.Updated_at,
			&LocaisObj.Delete_at,
			&LocaisObj.Nome,
			&LocaisObj.Endereco,
			&LocaisObj.Presencial,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Locais{}, err
		}
		LocaisList = append(LocaisList, LocaisObj)
	}
	rows.Close()
	return LocaisList, nil
}

func (us *LocalRepository) CreateLocais(local model.Locais) (int, error) {
	query, err := us.connection.Prepare("INSERT INTO tb_local (nome, presencial, endereco) VALUES (?, ?, ?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(local.Nome, local.Presencial, local.Endereco)
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
func (us *LocalRepository) DeleteLocais(locais model.Locais) (int, error) {
	query, err := us.connection.Prepare("UPDATE tb_local SET deleted_at = ? WHERE local_id = ?")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(locais.Delete_at, locais.Local_id)
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
