package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase/compromisso"
	"go-api/viewmodel"
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
func (u *CompromissoController) UpdateParticipanteCompromisso(ctx echo.Context) error {
	compromissoUUID := ctx.Param("uuid")

	var updateParticipantesCompromissos viewmodel.UpdateParticipantesCompromisso
	err := ctx.Bind(&updateParticipantesCompromissos)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	fmt.Printf("%v\n\n\n", updateParticipantesCompromissos)
	valid := updateParticipantesCompromissos.Validate()
	if !valid {
		ctx.JSON(http.StatusBadRequest, nil) // Arrumar para retornar erro
	}

	compromisso_atualizado, err := u.CompromissoUsecase.UpdateParticipantesCompromisso(compromissoUUID, updateParticipantesCompromissos.Emails)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	fmt.Printf("%v\n\n", compromisso_atualizado)
	return ctx.JSON(http.StatusCreated, compromisso_atualizado) // Parsear para viewmodel

}
