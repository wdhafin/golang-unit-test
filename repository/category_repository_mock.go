package repository

import (
	"github.com/stretchr/testify/mock"
	"github.com/wdhafin/golang-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (m *CategoryRepositoryMock) FindByID(ID string) *entity.Category {
	arguments := m.Mock.Called(ID)
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
