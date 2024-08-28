package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type CategoriasRepository struct {
	connection *sql.DB
}

func NewCategoriaRepository(connection *sql.DB) *CategoriasRepository {
	return &CategoriasRepository{
		connection: connection,
	}
}

func (us *CategoriasRepository) GetCategoriaRepository() ([]model.Categoria, error) {
	query := "SELECT categoria_id, created_at, updated_at, delete_at, nome, slug  FROM tb_categoria"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Categoria{}, err
	}
	var CategoriaList []model.Categoria
	var CategoriaObj model.Categoria

	for rows.Next() {
		err = rows.Scan(
			&CategoriaObj.Categoria_id,
			&CategoriaObj.Created_at,
			&CategoriaObj.Updated_at,
			&CategoriaObj.Delete_at,
			&CategoriaObj.Nome,
			&CategoriaObj.Slug,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Categoria{}, err
		}
		CategoriaList = append(CategoriaList, CategoriaObj)
	}
	rows.Close()
	return CategoriaList, nil
}

func (us *CategoriasRepository) CreateCategoria(categoria model.Categoria) (int, error) {
	query, err := us.connection.Prepare("INSERT INTO tb_categoria (nome, slug) VALUES (?,?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(categoria.Nome, categoria.Slug)
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
func (us *CategoriasRepository) DeleteCategoria(categoria model.Categoria) (int, error) {
	query, err := us.connection.Prepare("UPDATE tb_categoria SET deleted_at = ? WHERE categoria_id = ?")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(categoria.Delete_at, categoria.Categoria_id)
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
