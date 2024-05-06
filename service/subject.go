package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type subjectService struct {
	storage storage.IStorage
}

func NewSubjectService(storage storage.IStorage) subjectService {
	return subjectService{storage: storage}
}

func (t subjectService) CreateSubject(subjects models.Subjects) (string, error) {

	id, err := t.storage.SubjectStorage().CreateSubject(subjects)
	if err != nil {
		fmt.Println("error while creating teacher, err: ", err)
		return "", err
	}

	return id, nil
}

func (t subjectService) UpdateSubject(subject models.Subjects) error {
	err:=t.storage.SubjectStorage().UpadateSubject(subject)	
	if err!=nil {
		return err
	}
	return nil
}


func (t subjectService) GetbyIdSubject(id string) (models.Subjects,error) {
	resp,err:=t.storage.SubjectStorage().GetbyIdSubject(id)
	if err!=nil {
		return resp,err
	}

	return resp,nil
}

func (t subjectService) DeleteSubject(id string) error {
	err:=t.storage.SubjectStorage().DeleteSubject(id)
	if err!=nil {
		fmt.Println("error while deleting subject:",err)
		return err
	}

	return nil
}