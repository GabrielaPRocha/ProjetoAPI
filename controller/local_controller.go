package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LocaisController struct {
	LocalUseCase usecase.LocalUseCase
}

func NewLocaisController(usecase usecase.LocalUseCase) LocaisController {
	return LocaisController{
		LocalUseCase: usecase,
	}
}

func (u *LocaisController) GetLocal(ctx echo.Context) error {
	locais, err := u.LocalUseCase.GetLocais()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, locais)
}
func (u *LocaisController) CreateLocais(ctx echo.Context) error {
	var local model.Locais
	err := ctx.Bind(&local)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedLocal, err := u.LocalUseCase.CreateLocais(local)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, insertedLocal)
}
