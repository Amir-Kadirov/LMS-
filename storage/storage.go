package storage

import (
	"backend_course/lms/api/models"
	"context"
	"time"
)

type IStorage interface {
	CloseDB()
	StudentStorage() StudentStorage
	TeacherStorage() TeacherStorage
	SubjectStorage() SubjectStorage
	TimeTableStorage() TimeTableStorage
	Redis() IRedisStorage
}

type StudentStorage interface {
	Create(ctx context.Context,student models.Student) (string, error)
	GetAll(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error)
	UpdateSt(ctx context.Context,student models.Student) (string, error)
	GetById(ctx context.Context,id string) (models.GetStudent, error)
	DeleteSt(ctx context.Context,external_id string) error
	StatusSt(ctx context.Context,id string) (models.IsActiveResponse, error)
	CheckLessonStudent(ctx context.Context, id string) (models.CheckLessonStudent, error)
}

type TeacherStorage interface {
	CreateTeacher(ctx context.Context,teacher models.Teacher) (string, error)
	UpdateTeacher(ctx context.Context,teacher models.Teacher) (string,error)
	GetAllTeacher(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error)
	GetTeacherbyId(ctx context.Context,id string) (models.GetByIdTeacher, error)
	DeleteTeacher(ctx context.Context,id string) error
	CheckLessonTeacher(ctx context.Context, id string) (models.CheckLessonTeacher, error)
	GetTeacherbyLogin(ctx context.Context,login string) (models.Teacher, error)
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
	GetAllStudentsAttandenceReport(ctx context.Context, req models.GetAllStudentsAttandenceReportRequest) (models.GetAllStudentsAttandenceReportResponse, error)
}

type IRedisStorage interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	Get(ctx context.Context, key string) interface{}
	Del(ctx context.Context, key string) error
}
