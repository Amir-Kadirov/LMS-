package storage

import "backend_course/lms/api/models"

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
	SubjectStorage() SubjectStorage
	TimeTableStorage() TimeTableStorage
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
	UpdateTeacher(teacher models.Teacher) (string,error)
	GetAllTeacher(req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error)
	GetTeacherbyId(id string) (models.GetByIdTeacher, error)
	DeleteTeacher(id string) error
}

type SubjectStorage interface {
	CreateSubject(subject models.Subjects) (string, error)
	UpadateSubject(subject models.Subjects) (string,error)
	GetbyIdSubject(id string) (models.Subjects, error)
	DeleteSubject(id string) error
	GetAllSubject(req models.GetAllStudentsRequest) (models.SubjectGetAll, error)
}

type TimeTableStorage interface {
	DeleteTimeTable(id string) error
	CreateTimeTable(timeTable models.TimeTable) (string, error)
}
