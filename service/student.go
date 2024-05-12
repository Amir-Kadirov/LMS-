package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"fmt"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{storage: storage}
}

func (s studentService) Create(ctx context.Context,student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(ctx,student)
	if err != nil {
		fmt.Println("error while creating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) GetAllStudent(ctx context.Context,req models.GetAllStudentsRequest) (models.GetAllStudentsResponse,error) {
	students:=models.GetAllStudentsResponse{}

	students,err:=s.storage.StudentStorage().GetAll(ctx,req)
	if err!=nil {
		fmt.Println("error while get all students err: ",err)
		return students,err
	}

	return students,nil
}

func (s studentService) UpdateStudent(ctx context.Context,student models.Student) (string,error) {
	id,err:=s.storage.StudentStorage().UpdateSt(ctx,student)
	if err!=nil {
		return "",err
	}

	return id,nil
}

func (s studentService) GetByIdStudent(ctx context.Context,id string) (models.GetStudent,error) {
	student:=models.GetStudent{}
	student,err:=s.storage.StudentStorage().GetById(ctx,id)
	if err!=nil {
		fmt.Println("error while get by id student err: ",err)
		return student,err
	}

	return student,err
}

func (s studentService) DeleteStudent(ctx context.Context,id string) error {
	err:=s.storage.StudentStorage().DeleteSt(ctx,id)
	if err!=nil {
		fmt.Println("error while deleting student",err)
		return err
	}

	return nil
}

func (s studentService) StatusStudent(ctx context.Context,id string) (models.IsActiveResponse,error) {
	isactive,err:=s.storage.StudentStorage().StatusSt(ctx,id)
	if err!=nil {
		fmt.Println("error while checking status", err)
		return isactive,err
	}
	
	return isactive,nil
}

func (s studentService) LessonStudent(ctx context.Context,id string) (models.CheckLessonStudent,error) {
	lesson,err:=s.storage.StudentStorage().CheckLessonStudent(ctx,id)
	if err!=nil {
		return lesson,err
	}

	return lesson,nil
}