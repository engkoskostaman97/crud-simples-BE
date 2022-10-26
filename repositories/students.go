package repositories

import (
	"pelajar/models"

	"gorm.io/gorm"
)

type StudentRepository interface {
	FindStudents() ([]models.Student, error)
	GetStudent(ID int) (models.Student, error)
	CreateStudent(student models.Student) (models.Student, error)
	UpdateStudent(student models.Student) (models.Student, error)
	DeleteStudent(student models.Student) (models.Student, error)
}

type repository struct {
	db *gorm.DB
}

func RepositoryStudent(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindStudents() ([]models.Student, error) {
	var students []models.Student
	err := r.db.Find(&students).Error

	return students, err
}

func (r *repository) GetStudent(ID int) (models.Student, error) {
	var students models.Student
	err := r.db.First(&students	, ID).Error

	return students, err
}

func (r *repository) CreateStudent(student models.Student) (models.Student, error) {
	err := r.db.Create(&student).Error

	return student, err
}

func (r *repository) UpdateStudent(student models.Student) (models.Student, error) {
	err := r.db.Save(&student).Error

	return student, err
}

func (r *repository) DeleteStudent(student models.Student) (models.Student, error) {

	err := r.db.Delete(&student).Error

	return student, err
}
