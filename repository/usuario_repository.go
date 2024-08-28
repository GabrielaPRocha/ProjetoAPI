package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"

	uuid "github.com/google/uuid"
)

type UsuarioRepository struct {
	connection *sql.DB
}

func NewUsuarioRepository(connection *sql.DB) UsuarioRepository {
	return UsuarioRepository{
		connection: connection,
	}
}
func (us *UsuarioRepository) GetUsuarios() ([]model.Usuarios, error) {
	query := "SELECT usuario_id,created_at,updated_at,delete_at,nome,email,status,senha,uuid FROM tb_usuarios"
	rows, err := us.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Usuarios{}, err
	}
	var usuarioList []model.Usuarios
	var usuariosObj model.Usuarios

	for rows.Next() {
		err = rows.Scan(
			&usuariosObj.Usuario_id,
			&usuariosObj.Created_at,
			&usuariosObj.Updated_at,
			&usuariosObj.Delete_at,
			&usuariosObj.Nome,
			&usuariosObj.Email,
			&usuariosObj.Status,
			&usuariosObj.Senha,
			&usuariosObj.Uuid_usuario,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Usuarios{}, err
		}
		usuarioList = append(usuarioList, usuariosObj)
	}
	rows.Close()
	return usuarioList, nil

}
func (us *UsuarioRepository) CreateUsuarios(usuario model.Usuarios) (int, error) {

	usuario.Uuid_usuario = uuid.New().String()
	query, err := us.connection.Prepare("INSERT INTO tb_usuarios" +
		"(nome,email,status,senha,uuid)" +
		"VALUES(?, ?, ?, ?,?)")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	defer query.Close()

	result, err := query.Exec(usuario.Nome, usuario.Email, usuario.Status, usuario.Senha, usuario.Uuid_usuario)
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
