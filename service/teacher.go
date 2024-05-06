package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type teacherService struct {
	storage storage.IStorage
}

func NewTeacherService(storage storage.IStorage) teacherService {
	return teacherService{storage: storage}
}

func (t teacherService) CreateTeacher(teacher models.Teacher) (string, error) {

	id, err := t.storage.TeacherStorage().CreateTeacher(teacher)
	if err != nil {
		fmt.Println("error while creating teacher, err: ", err)
		return "", err
	}

	return id, nil
}

func (t teacherService) UpdateTeacher(teacher models.Teacher) error {
	err := t.storage.TeacherStorage().UpdateTeacher(teacher)
	if err != nil {
		fmt.Println("error while updating teacher:", err)
		return err
	}
	return nil
}

func (t teacherService) GetAllTeacher(req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error) {

	teachers,err:=t.storage.TeacherStorage().GetAllTeacher(req)
	if err!=nil {
		fmt.Println("error while get all teacher:",err)
		return teachers,err
	}

	return teachers,nil
}

func (t teacherService) GetTeacherbyId(id string) (models.Teacher, error) {
	teacher,err:=t.storage.TeacherStorage().GetTeacherbyId(id)
	if err!=nil {
		fmt.Println("error while get by id teacher:",err)
		return teacher,err
	}

	return teacher,nil
}

func (t teacherService) DeleteTeacher(id string) error {
	err:=t.storage.TeacherStorage().DeleteTeacher(id)
	if err!=nil {
		fmt.Println("error while deleting teacher:",err)
		return err
	}

	return nil
}