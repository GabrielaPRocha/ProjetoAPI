package controller

import (
	"go-api/model"
	"go-api/usecase/usuario"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsuariosController struct {
	UsuarioUsecase usuario.UsuarioUsecase
}

func NewUsuariosController(usecase usuario.UsuarioUsecase) UsuariosController {
	return UsuariosController{
		UsuarioUsecase: usecase,
	}
}

func (u *UsuariosController) GetUsuario(ctx echo.Context) error {
	usuarios, err := u.UsuarioUsecase.GetUsuarios()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, usuarios)
}
func (u *UsuariosController) CreateUsuarios(ctx echo.Context) error {
	var usuario model.Usuarios
	err := ctx.Bind(&usuario)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedUsuario, err := u.UsuarioUsecase.CreateUsuarios(usuario)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, insertedUsuario)
}
