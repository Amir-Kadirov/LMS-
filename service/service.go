package service

import "backend_course/lms/storage"

type IServiceManager interface {
	Student() studentService
	Teacher() teacherService
	Subject() subjectService
	TimeTable() timetableService
}

type Service struct {
	studentService studentService
	teacherService teacherService
	subjectService subjectService
	timetableService timetableService
}

func New(storage storage.IStorage) Service {
	services := Service{}
	services.studentService = NewStudentService(storage)
	services.teacherService = NewTeacherService(storage)
	services.subjectService = NewSubjectService(storage)
	services.timetableService=NewTimeTableService(storage)

	return services
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