package repository

import (
	"github.com/Zeta201/examination-service/infrastructure"
	"github.com/Zeta201/examination-service/model"
)

type DepartmentRepository struct {
	db infrastructure.Database
}

func NewDepartmentRepository(db infrastructure.Database) DepartmentRepository {
	return DepartmentRepository{
		db: db,
	}
}
func (repository DepartmentRepository) SaveDepartment(dept model.Department) error {
	return repository.db.DB.Create(&dept).Error
}
