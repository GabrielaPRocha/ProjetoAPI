package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CompromissoController struct {
	CompromissoUsecase usecase.CompromissoUsecase
}

func NewCompromissoController(usecase usecase.CompromissoUsecase) CompromissoController {
	return CompromissoController{
		CompromissoUsecase: usecase,
	}
}

func (u *CompromissoController) GetCompromisso(ctx echo.Context) error {
	compromissos, err := u.CompromissoUsecase.GetCompromissos()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
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
