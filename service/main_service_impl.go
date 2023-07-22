package service

import (
	"perpustakaan/entity"
	"perpustakaan/repository"
)

type MainServiceImpl struct {
	Repository repository.MainRepository
}

func NewMainServiceImpl(mainRepository repository.MainRepository) *MainServiceImpl {
	return &MainServiceImpl{
		Repository: mainRepository,
	}
}

func (service *MainServiceImpl) FindAll() entity.Data {
	//TODO implement me
	books := service.Repository.FindAll()
	return books
}
