package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase/categoria"
	"go-api/usecase/compromisso"
	"go-api/usecase/local"
	"go-api/usecase/participante"
	"go-api/usecase/usuario"

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
	UsuariosUseCase := usuario.NewUsuariosUseCase(UsuarioRepository)
	UsuariosController := controller.NewUsuariosController(UsuariosUseCase)
	//Local
	LocalRepository := repository.NewLocalRepository(dbConnection)
	LocalUseCase := local.NewLocalUseCase(LocalRepository)
	LocaisController := controller.NewLocaisController(LocalUseCase)
	//Categoria
	CategoriaRepository := repository.NewCategoriaRepository(dbConnection)
	CategoriaUseCase := categoria.NewCategoriaUseCase(CategoriaRepository)
	CategoriaController := controller.NewCategoriaController(CategoriaUseCase)

	//Participante

	ParticipanteRepository := repository.NewParticipanteRepository(dbConnection)
	ParticipantesUseCase := participante.NewParticipanteUsecase(*ParticipanteRepository)
	ParticipantesController := controller.NewParticipantesController(ParticipantesUseCase)

	CompromissoRepository := repository.NewCompromissoRepository(dbConnection)
	CompromissoUsecase := compromisso.NewCompromissoUsecase(*CompromissoRepository, ParticipantesUseCase)
	CompromissoController := controller.NewCompromissoController(CompromissoUsecase)

	server.GET("/usuarios", UsuariosController.GetUsuario)
	server.POST("/criar", UsuariosController.CreateUsuarios)
	server.GET("/local", LocaisController.GetLocal)
	server.POST("/criarlocal", LocaisController.CreateLocais)
	server.GET("/categoria", CategoriaController.GetCategoria)
	server.POST("/criarCategoria", CategoriaController.CreateCategoria)
	server.GET("/participante", ParticipantesController.GetParticipantes)
	server.POST("/criarParticipante", ParticipantesController.CreateParticipante)
	server.GET("/compromissos/:uuid", CompromissoController.GetCompromisso)
	server.POST("/criarCompromisso", CompromissoController.CreateCompromisso)
	server.PATCH("/compromissos/:uuid/participantes", CompromissoController.UpdateParticipanteCompromisso)

	server.Logger.Fatal(server.Start(":8080"))

}
