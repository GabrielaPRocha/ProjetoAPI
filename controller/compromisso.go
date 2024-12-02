package controller

import (
	"fmt"
	"go-api/model"
	"go-api/usecase/compromisso"
	"net/http"
	"regexp"

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
	usuarioUUID := ctx.Param("uuid")
	email, err := u.CompromissoUsecase.GetCompromissos(usuarioUUID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	var updatecompromissos model.UpdateCompromisso
	fmt.Printf("%v\n\n\n", updatecompromissos)

	err = ctx.Bind(&updatecompromissos)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}
	updatecompromissos = email

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	nameRegex := `^[a-zA-Z\s]+$`
	emailPattern, _ := regexp.Compile(emailRegex)
	namePattern, _ := regexp.Compile(nameRegex)

	if !emailPattern.MatchString(updatecompromissos.Email) {
		return ctx.JSON(http.StatusBadRequest, err)
	}
	if !namePattern.MatchString(updatecompromissos.Nome) {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	fmt.Printf("%v\n\n\n\n", updatecompromissos)
	upCompromisso, err := u.CompromissoUsecase.GetCompromissos(updatecompromissos)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	fmt.Printf("%v\n\n", upCompromisso)
	return ctx.JSON(http.StatusCreated, upCompromisso)

}
