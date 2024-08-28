package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CategoriaController struct {
	CategoriaUseCase usecase.CategoriaUseCase
}

func NewCategoriaController(usecase usecase.CategoriaUseCase) CategoriaController {
	return CategoriaController{
		CategoriaUseCase: usecase,
	}
}

func (u *CategoriaController) GetCategoria(ctx echo.Context) error {
	categorias, err := u.CategoriaUseCase.GetCategoria()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, categorias)
}
func (u *CategoriaController) CreateCategoria(ctx echo.Context) error {
	var categoria model.Categoria
	err := ctx.Bind(&categoria)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedCategoria, err := u.CategoriaUseCase.CreateCategoria(categoria)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, insertedCategoria)
}
func (u *CategoriaController) DeleteCategoria(ctx echo.Context) error {
	var categoria model.Categoria
	err := ctx.Bind(&categoria)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	deletedCategoria, err := u.CategoriaUseCase.DeleteCategoria(categoria)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, deletedCategoria)
}
