package service

import (
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
)

type IServiceManager interface {
	Student() studentService
	Teacher() teacherService
	Subject() subjectService
	TimeTable() timetableService
	Auth() authService
}

type Service struct {
	studentService studentService
	teacherService teacherService
	subjectService subjectService
	timetableService timetableService
	authService authService

	logger logger.ILogger
}

func New(storage storage.IStorage,log logger.ILogger) Service {
	return Service{
		studentService: NewStudentService(storage,log),
		subjectService: NewSubjectService(storage,log),
		teacherService: NewTeacherService(storage,log),
		timetableService: NewTimeTableService(storage,log),
		authService: NewAuthService(storage,log),

		logger: log,
	}
}

func (s Service) Student() studentService {
	return s.studentService
}

func (s Service) Teacher() teacherService {
	return s.teacherService
}

func (s Service) Subject() subjectService {
	return s.subjectService
}

func (s Service) TimeTable() timetableService {
	return s.timetableService
}

func (s Service) Auth() authService {
	return s.authService
}