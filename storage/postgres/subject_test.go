package postgres

import (
	"backend_course/lms/api/models"
	"context"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateSubject(t *testing.T) {
	SubjectRepo := NewSubject(db)

	reqSubject := models.Subjects{
		Name: faker.Name(),
		Type: faker.Paragraph(),
	}

	id, err := SubjectRepo.CreateSubject(context.Background(), reqSubject)
	if assert.NoError(t, err) {
		createdSubject, err := SubjectRepo.GetbyIdSubject(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqSubject.Name, createdSubject.Name)
			assert.Equal(t, reqSubject.Type, createdSubject.Type)

		} else {
			return
		}
		fmt.Println("Created Subject", createdSubject)
	}
}

func TestUpdateSubject(t *testing.T) {
	SubjectRepo := NewSubject(db)
	updateSubject := models.Subjects{
		Name: faker.Name(),
		Type: faker.Paragraph(),
		Id:   "efc61203-c724-4c56-9eef-0c219a8b7ecb",
	}
	id, err := SubjectRepo.UpadateSubject(context.Background(), updateSubject)
	if assert.NoError(t, err) {
		updatedSubject, err := SubjectRepo.GetbyIdSubject(context.Background(), id)
		if assert.NoError(t, err) {
			assert.Equal(t, updateSubject.Name, updatedSubject.Name)
			assert.Equal(t, updateSubject.Type, updatedSubject.Type)
		}
	}
}

func TestDeleteSubject(t *testing.T) {
	SubjectRepo := NewSubject(db)

	id := "752a60b8-e934-4bb3-9bf0-45040acfa208"

	erDr := SubjectRepo.DeleteSubject(context.Background(), id)
	if assert.NoError(t, erDr) {
		_, err := SubjectRepo.GetbyIdSubject(context.Background(), id)
		assert.Error(t, err, err)
	}
}

func TestGetAllSubject(t *testing.T) {
	SubjectRepo := NewSubject(db)
	reqSubject := models.Subjects{
		Name: faker.Name(),
		Type: faker.Word(),
	}

	getAllreq := models.GetAllStudentsRequest{
		Search: "",
		Page:   1,
		Limit:  10,
	}

	_, err := SubjectRepo.CreateSubject(context.Background(), reqSubject)
	oldCount, _ := SubjectRepo.GetAllSubject(context.Background(), getAllreq)

	if assert.NoError(t, err) {
		_, err = SubjectRepo.CreateSubject(context.Background(), reqSubject)
		if assert.NoError(t, err) {
			newcount, _ := SubjectRepo.GetAllSubject(context.Background(), getAllreq)
			assert.Equal(t, 0, newcount.Count-oldCount.Count)
		}
	}
}
