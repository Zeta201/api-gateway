package service

import (
	"github.com/Zeta201/examination-service/model"
	"github.com/Zeta201/examination-service/repository"
)

type SubjectService struct {
	repository repository.SubjectRepository
}

func NewSubjectService(r repository.SubjectRepository) SubjectService {
	return SubjectService{
		repository: r,
	}
}

func (service SubjectService) SaveSubject(subject model.Subject) error {
	return service.repository.SaveSubject(subject)
}

func (service SubjectService) FindAll(subject model.Subject, keyword string) (*[]model.Subject, int64, error) {
	return service.repository.GetAllSubjects(subject, keyword)
}

// func (service StudentService) Update(student model.Student) error{
// 	return service.repository.Update(student)
// }

// func (service StudentService) Delete(id int64) error {
// 	var student model.Student
// 	student.ID = id
// 	return service.repository.Delete(student)
// }

func (service SubjectService) GetSubjectById(subject model.Subject) (model.Subject, error) {
	return service.repository.GetSubjectById(subject)
}
