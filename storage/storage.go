package storage

import (
	"backend_course/lms/api/models"
	"context"
)

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
	SubjectStorage() SubjectStorage
	TimeTableStorage() TimeTableStorage
}

type StudentStorage interface {
	Create(ctx context.Context,student models.Student) (string, error)
	GetAll(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	UpdateSt(ctx context.Context,student models.Student) (string, error)
	UpdateStPassword(ctx context.Context,id string, password string) (string, error)
	GetById(ctx context.Context,ExternalId string) (models.GetStudent, error)
	DeleteSt(ctx context.Context,external_id string) error
	StatusSt(ctx context.Context,id string) (models.IsActiveResponse, error)
}

type TeacherStorage interface {
	CreateTeacher(ctx context.Context,teacher models.Teacher) (string, error)
	UpdateTeacher(ctx context.Context,teacher models.Teacher) (string,error)
	GetAllTeacher(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error)
	GetTeacherbyId(ctx context.Context,id string) (models.GetByIdTeacher, error)
	DeleteTeacher(ctx context.Context,id string) error
}

type SubjectStorage interface {
	CreateSubject(ctx context.Context,subject models.Subjects) (string, error)
	UpadateSubject(ctx context.Context,subject models.Subjects) (string,error)
	GetbyIdSubject(ctx context.Context,id string) (models.Subjects, error)
	DeleteSubject(ctx context.Context,id string) error
	GetAllSubject(ctx context.Context,req models.GetAllStudentsRequest) (models.SubjectGetAll, error)
}

type TimeTableStorage interface {
	DeleteTimeTable(ctx context.Context,id string) error
	CreateTimeTable(ctx context.Context,timeTable models.TimeTable) (string, error)
}
