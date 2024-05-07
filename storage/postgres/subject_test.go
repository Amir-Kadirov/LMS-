package postgres

import (
	"backend_course/lms/api/models"
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

	id, err := SubjectRepo.CreateSubject(reqSubject)
	if assert.NoError(t, err) {
		createdSubject, err := SubjectRepo.GetbyIdSubject(id)
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
		Id:        "3ed68793-cd2a-4c0a-9358-0eb2faa05cf5",
	}
	id, err := SubjectRepo.UpadateSubject(updateSubject)
	if assert.NoError(t, err) {
		updatedSubject, err := SubjectRepo.GetbyIdSubject(id)
		if assert.NoError(t, err) {
			assert.Equal(t, updateSubject.Name, updatedSubject.Name)
			assert.Equal(t, updateSubject.Type, updatedSubject.Type)
		}
	}
}

func TestDeleteSubject(t *testing.T) {
	SubjectRepo := NewSubject(db)

	id := "a0e27142-e55e-44ab-b292-5880f79b4243"

	erDr := SubjectRepo.DeleteSubject(id)
	if assert.NoError(t, erDr) {
		_, err := SubjectRepo.GetbyIdSubject(id)
		assert.Error(t, err, err)
	}
}

func TestGetAllSubject(t *testing.T) {
	SubjectRepo := NewSubject(db)
	reqSubject := models.Subjects{
		Name: faker.Name(),
		Type:  faker.Word(),
	}

	getAllreq := models.GetAllStudentsRequest{
		Search: "",
		Page:   1,
		Limit:  10,
	}

	_, err := SubjectRepo.CreateSubject(reqSubject)
	oldCount, _ := SubjectRepo.GetAllSubject(getAllreq)

	if assert.NoError(t, err) {
		_, err = SubjectRepo.CreateSubject(reqSubject)
		if assert.NoError(t, err) {
			newcount, _ := SubjectRepo.GetAllSubject(getAllreq)
			assert.Equal(t, 1, newcount.Count-oldCount.Count)
		}
	}
}