package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/wdhafin/golang-unit-test/entity"
	"github.com/wdhafin/golang-unit-test/repository"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryServiceGetNotFound(t *testing.T) {

	// program mock
	categoryRepository.Mock.On("FindByID", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryServiceGetSuccess(t *testing.T) {
	category := entity.Category{
		ID:   "2",
		Name: "Laptop",
	}
	categoryRepository.Mock.On("FindByID", "2").Return(category)
	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}
