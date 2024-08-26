package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	UsuarioRepository := repository.NewUsuarioRepository(dbConnection)
	UsuariosUseCase := usecase.NewUsuariosUseCase(UsuarioRepository)
	UsuariosController := controller.NewUsuariosController(UsuariosUseCase)

	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "Hello World",
		})
	})
	server.GET("/usuarios", UsuariosController.GetUsuario)
	server.POST("/criar", UsuariosController.CreateUsuarios)
	server.Logger.Fatal(server.Start(":8081"))

}
