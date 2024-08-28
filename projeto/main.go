package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	server := echo.New()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}
	//Usuario
	UsuarioRepository := repository.NewUsuarioRepository(dbConnection)
	UsuariosUseCase := usecase.NewUsuariosUseCase(UsuarioRepository)
	UsuariosController := controller.NewUsuariosController(UsuariosUseCase)
	//Local
	LocalRepository := repository.NewLocalRepository(dbConnection)
	LocalUseCase := usecase.NewLocalUseCase(LocalRepository)
	LocaisController := controller.NewLocaisController(LocalUseCase)
	//Categoria
	CategoriaRepository := repository.NewCategoriaRepository(dbConnection)
	CategoriaUseCase := usecase.NewCategoriaUseCase(CategoriaRepository)
	CategoriaController := controller.NewCategoriaController(CategoriaUseCase)

	server.GET("/usuarios", UsuariosController.GetUsuario)
	server.POST("/criar", UsuariosController.CreateUsuarios)
	server.GET("/local", LocaisController.GetLocal)
	server.POST("/criarlocal", LocaisController.CreateLocais)
	server.GET("/categoria", CategoriaController.GetCategoria)
	server.POST("/criarCategoria", CategoriaController.CreateCategoria)
	server.Logger.Fatal(server.Start(":8081"))

}
