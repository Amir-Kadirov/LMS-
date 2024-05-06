package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type studentService struct {
	storage storage.IStorage
}

func NewStudentService(storage storage.IStorage) studentService {
	return studentService{storage: storage}
}

func (s studentService) Create(student models.Student) (string, error) {
	// business logic
	id, err := s.storage.StudentStorage().Create(student)
	if err != nil {
		fmt.Println("error while creating student, err: ", err)
		return "", err
	}
	// business logic
	return id, nil
}

func (s studentService) GetAllStudent(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse,error) {
	students:=models.GetAllStudentsResponse{}

	students,err:=s.storage.StudentStorage().GetAll(req)
	if err!=nil {
		fmt.Println("error while get all students err: ",err)
		return students,err
	}

	return students,nil
}

func (s studentService) UpdateStudent(student models.Student) (string,error) {
	id,err:=s.storage.StudentStorage().UpdateSt(student)
	if err!=nil {
		return "",err
	}

	return id,nil
}

func (s studentService) GetByIdStudent(id string) (models.GetStudent,error) {
	student:=models.GetStudent{}
	student,err:=s.storage.StudentStorage().GetById(id)
	if err!=nil {
		fmt.Println("error while get by id student err: ",err)
		return student,err
	}

	return student,err
}

func (s studentService) DeleteStudent(id string) error {
	err:=s.storage.StudentStorage().DeleteSt(id)
	if err!=nil {
		fmt.Println("error while deleting student",err)
		return err
	}

	return nil
}

func (s studentService) UpdatePassword(id string,password string) (string,error) {
	password,err:=s.storage.StudentStorage().UpdateStPassword(id,password)
	if err!=nil {
		fmt.Println("error while updating password ",err)
		return "",err
	}

	return password,nil
}