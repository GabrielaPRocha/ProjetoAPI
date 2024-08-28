package usecase

import (
	"go-api/model"
	"go-api/repository"
)

type CategoriaUseCase struct {
	repository repository.CategoriasRepository
}

func NewCategoriaUseCase(repo *repository.CategoriasRepository) CategoriaUseCase {
	return CategoriaUseCase{
		repository: *repo,
	}
}

func (su *CategoriaUseCase) GetCategoria() ([]model.Categoria, error) {
	return su.repository.GetCategoriaRepository()
}

func (us *CategoriaUseCase) CreateCategoria(Categoria model.Categoria) (model.Categoria, error) {
	CategoriaId, err := us.repository.CreateCategoria(Categoria)
	if err != nil {
		return model.Categoria{}, err
	}
	Categoria.Categoria_id = CategoriaId
	return Categoria, nil
}
func (us *CategoriaUseCase) DeleteCategoria(categoria model.Categoria) (model.Categoria, error) {
	_, err := us.repository.DeleteCategoria(categoria)
	if err != nil {
		return model.Categoria{}, err
	}
	return categoria, nil
}
