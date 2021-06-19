package service

import (
	"errors"

	"github.com/wdhafin/golang-unit-test/entity"
	"github.com/wdhafin/golang-unit-test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(ID string) (*entity.Category, error) {
	category := service.Repository.FindByID(ID)
	if category == nil {
		return nil, errors.New("category not found")
	} else {
		return category, nil
	}
}
