package service

import (
	"github.com/Zeta201/examination-service/model"
	"github.com/Zeta201/examination-service/repository"
)

type DepartmentService struct {
	repository repository.DepartmentRepository
}

func NewDepartmentService(r repository.DepartmentRepository) DepartmentService {
	return DepartmentService{
		repository: r,
	}
}

func (service DepartmentService) SaveDepartment(dept model.Department) error {
	return service.repository.SaveDepartment(dept)
}
