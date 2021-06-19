package repository

import "github.com/wdhafin/golang-unit-test/entity"

type CategoryRepository interface {
	FindByID(ID string) *entity.Category
}
