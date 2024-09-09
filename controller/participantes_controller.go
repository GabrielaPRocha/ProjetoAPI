package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ParticipantesController struct {
	ParticipanteUsecase usecase.ParticipanteUsecase
}

func NewParticipantesController(usecase usecase.ParticipanteUsecase) ParticipantesController {
	return ParticipantesController{
		ParticipanteUsecase: usecase,
	}
}

func (u *ParticipantesController) GetParticipantes(ctx echo.Context) error {
	participantes, err := u.ParticipanteUsecase.GetParticipantes()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, participantes)
}
func (u *ParticipantesController) CreateParticipante(ctx echo.Context) error {
	var participantes model.Participantes
	err := ctx.Bind(&participantes)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	}

	insertedParticipantes, err := u.ParticipanteUsecase.CreateParticipante(participantes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusCreated, insertedParticipantes)
}
func (u *ParticipantesController) DeleteParticipantes(ctx echo.Context) error {
	var participantes model.Participantes
	err := ctx.Bind(&participantes)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	deletedParticipantes, err := u.ParticipanteUsecase.DeleteParticipantes(participantes)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, deletedParticipantes)
}
