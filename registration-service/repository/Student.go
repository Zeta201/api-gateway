package repository

import (
	"github.com/Zeta201/registration-service/infrastructure"
	"github.com/Zeta201/registration-service/model"
)

type StudentRepository struct {
	db infrastructure.Database
}

func NewStudentRepository(db infrastructure.Database) StudentRepository {
	return StudentRepository{
		db: db,
	}
}

func (repository StudentRepository) Save(student model.Student) error {
	return repository.db.DB.Create(&student).Error
}

func (repository StudentRepository) FindAll(student model.Student, keyword string) (*[]model.Student, int64, error) {
	var students []model.Student
	var totalRows int64 = 0

	queryBuilder := repository.db.DB.Order("created_at desc").Model(&model.Student{})
	if keyword != "" {
		queryKeyword := "%" + keyword + "%"
		queryBuilder = queryBuilder.Where("Student.firstname like?", queryKeyword)
	}

	err := queryBuilder.Where(student).Find(&students).Count(&totalRows).Error
	return &students, totalRows, err
}

func (repository StudentRepository) Find(student model.Student) (model.Student, error) {
	var students model.Student
	err := repository.db.DB.Debug().Model(&model.Student{}).Where(&student).Take(&students).Error
	return students, err
}

func (repository StudentRepository) Delete(student model.Student) error {
	return repository.db.DB.Delete(&student).Error
}
