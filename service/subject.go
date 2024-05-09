package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"fmt"
)

type subjectService struct {
	storage storage.IStorage
}

func NewSubjectService(storage storage.IStorage) subjectService {
	return subjectService{storage: storage}
}

func (t subjectService) CreateSubject(ctx context.Context,subjects models.Subjects) (string, error) {

	id, err := t.storage.SubjectStorage().CreateSubject(ctx,subjects)
	if err != nil {
		fmt.Println("error while creating teacher, err: ", err)
		return "", err
	}

	return id, nil
}

func (t subjectService) UpdateSubject(ctx context.Context,subject models.Subjects) error {
	_,err:=t.storage.SubjectStorage().UpadateSubject(ctx,subject)	
	if err!=nil {
		return err
	}
	return nil
}


func (t subjectService) GetbyIdSubject(ctx context.Context,id string) (models.Subjects,error) {
	resp,err:=t.storage.SubjectStorage().GetbyIdSubject(ctx,id)
	if err!=nil {
		return resp,err
	}

	return resp,nil
}

func (t subjectService) DeleteSubject(ctx context.Context,id string) error {
	err:=t.storage.SubjectStorage().DeleteSubject(ctx,id)
	if err!=nil {
		fmt.Println("error while deleting subject:",err)
		return err
	}

	return nil
}

func (t subjectService) GetAllSubject(ctx context.Context,req models.GetAllStudentsRequest) (models.SubjectGetAll, error) {

	subject,err:=t.storage.SubjectStorage().GetAllSubject(ctx,req)
	if err!=nil {
		fmt.Println("error while get all subject:",err)
		return subject,err
	}

	return subject,nil
}