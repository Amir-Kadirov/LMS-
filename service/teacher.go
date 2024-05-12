package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"fmt"
)

type teacherService struct {
	storage storage.IStorage
}

func NewTeacherService(storage storage.IStorage) teacherService {
	return teacherService{storage: storage}
}

func (t teacherService) CreateTeacher(ctx context.Context,teacher models.Teacher) (string, error) {

	id, err := t.storage.TeacherStorage().CreateTeacher(ctx,teacher)
	if err != nil {
		fmt.Println("error while creating teacher, err: ", err)
		return "", err
	}

	return id, nil
}

func (t teacherService) UpdateTeacher(ctx context.Context,teacher models.Teacher) error {
	_,err := t.storage.TeacherStorage().UpdateTeacher(ctx,teacher)
	if err != nil {
		fmt.Println("error while updating teacher:", err)
		return err
	}
	return nil
}

func (t teacherService) GetAllTeacher(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error) {

	teachers,err:=t.storage.TeacherStorage().GetAllTeacher(ctx,req)
	if err!=nil {
		fmt.Println("error while get all teacher:",err)
		return teachers,err
	}

	return teachers,nil
}

func (t teacherService) GetTeacherbyId(ctx context.Context,id string) (models.GetByIdTeacher, error) {
	teacher,err:=t.storage.TeacherStorage().GetTeacherbyId(ctx,id)
	if err!=nil {
		fmt.Println("error while get by id teacher:",err)
		return teacher,err
	}

	return teacher,nil
}

func (t teacherService) DeleteTeacher(ctx context.Context,id string) error {
	err:=t.storage.TeacherStorage().DeleteTeacher(ctx,id)
	if err!=nil {
		fmt.Println("error while deleting teacher:",err)
		return err
	}

	return nil
}

func (t teacherService) LessonTeacher(ctx context.Context,id string) (models.CheckLessonTeacher,error) {
	lesson,err:=t.storage.TeacherStorage().CheckLessonTeacher(ctx,id)
	if err!=nil {
		return lesson,err
	}

	return lesson,nil
}