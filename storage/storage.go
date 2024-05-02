package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
	GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	UpdateSt(student models.Student) (string, error)
	UpdateStPassword(student models.Student) (string, error)
	GetById(student models.GetByIdRequest) (models.GetStudent, error)
	DeleteSt(student models.Student) (string, error)
}