package main

import (
	"go-api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()
	UsuariosController := controller.NewUsuariosController()

	server.GET("/ping", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, map[string]string{
			"message": "Hello World",
		})
	})

	server.Logger.Fatal(server.Start(":8081"))
}
