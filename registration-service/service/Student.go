package service

import (
	"github.com/Zeta201/registration-service/model"
	"github.com/Zeta201/registration-service/repository"
)

type StudentService struct {
	repository repository.StudentRepository
}

func NewStudentService(r repository.StudentRepository) StudentService {
	return StudentService{
		repository: r,
	}
}

func (service StudentService) Save(student model.Student) error {
	return service.repository.Save(student)
}

func (service StudentService) FindAll(student model.Student, keyword string) (*[]model.Student, int64, error) {
	return service.repository.FindAll(student, keyword)
}

// func (service StudentService) Update(student model.Student) error{
// 	return service.repository.Update(student)
// }

func (service StudentService) Delete(id int64) error {
	var student model.Student
	student.ID = id
	return service.repository.Delete(student)
}

func (service StudentService) Find(student model.Student) (model.Student, error) {
	return service.repository.Find(student)
}
