package repository

import (
	"github.com/Zeta201/examination-service/infrastructure"
	"github.com/Zeta201/examination-service/model"
)

type SubjectRepository struct {
	db infrastructure.Database
}

func NewSubjectRepository(db infrastructure.Database) SubjectRepository {
	return SubjectRepository{
		db: db,
	}
}

func (repository SubjectRepository) SaveSubject(subject model.Subject) error {
	return repository.db.DB.Create(&subject).Error
}

func (repository SubjectRepository) GetAllSubjects(subject model.Subject, keyword string) (*[]model.Subject, int64, error) {
	var subjects []model.Subject
	var totalRows int64 = 0

	queryBuilder := repository.db.DB.Model(&model.Subject{})
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where("Subject.name like?", queryKeyword)
	}

	err := queryBuilder.Where(subject).Preload("Departments").Find(&subjects).Count(&totalRows).Error
	return &subjects, totalRows, err
}

func (repository SubjectRepository) GetSubjectById(subject model.Subject) (model.Subject, error) {
	var subjects model.Subject
	err := repository.db.DB.Debug().Model(&model.Subject{}).Where(&subject).Preload("Departments").Take(&subjects).Error
	return subjects, err
}

// func (repository SubjectRepository) Delete(student model.Student) error {
// 	return repository.db.DB.Delete(&student).Error
// }
