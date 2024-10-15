package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase/compromisso"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompromissoController struct {
	CompromissoUsecase compromisso.CompromissoUsecase
}

func NewCompromissoController(usecase compromisso.CompromissoUsecase) CompromissoController {
	return CompromissoController{
		CompromissoUsecase: usecase,
	}
}

func (u *CompromissoController) GetCompromisso(ctx echo.Context) error {
	usuarioUUID := ctx.Param("uuid")
	fmt.Printf("%v", usuarioUUID)
	compromissos, err := u.CompromissoUsecase.GetCompromissos(usuarioUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, compromissos)
}

func (u *CompromissoController) CreateCompromisso(ctx echo.Context) error {
	var compromissos model.Compromisso
	err := ctx.Bind(&compromissos)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedCompromisso, err := u.CompromissoUsecase.CreateCompromisso(compromissos)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, insertedCompromisso)

}
