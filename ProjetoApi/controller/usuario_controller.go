package controller

import (
	"ProjetoAPI/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsuariosController struct {
	//usecase
}

func NewUsuariosController() UsuariosController {
	return UsuariosController{}
}

func (p NewUsuariosController) GetUsuarios(ctx echo.Context) error {
	usuarios := []model.Usuarios{
		{
			usuario_id: 1,
			created_at: "1/1/1",
			updated_at: "1/1/1",
			delete_at:  "1/1/1",
			nome:       "Gabriela",
			email:      "gpereira@redventures.com",
			status:     true,
			Senha:      "123",
		},
	}
	return ctx.JSON(http.StatusOK, usuarios)
}
