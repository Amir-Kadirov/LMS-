package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
}

type StudentStorage interface {
	Create(student models.Student) (string, error)
	GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	UpdateSt(student models.Student) (string, error)
	UpdateStPassword(id string, password string) (string, error)
	GetById(ExternalId string) (models.GetStudent, error)
	DeleteSt(external_id string) error
	StatusSt(id string) (models.IsActiveResponse, error)
}

type TeacherStorage interface {
	CreateTeacher(teacher models.Teacher) (string, error)
	UpdateTeacher(teacher models.Teacher) error
	GetAllTeacher(req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error)
	GetTeacherbyId(id string) (models.Teacher, error)
	DeleteTeacher(id string) error
}