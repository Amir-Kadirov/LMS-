package postgres

import (
	"backend_course/lms/api/models"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	reqStudent := models.Student{
		FirstName: faker.Name(),
		Age:       10,
		LastName:  faker.Word(),
	}

	id, err := studentRepo.Create(reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetById(id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqStudent.FirstName, createdStudent.FirstName)
			assert.Equal(t, reqStudent.Age, createdStudent.Age)
			assert.Equal(t, reqStudent.LastName, createdStudent.LastName)

		} else {
			return
		}
		fmt.Println("Created student", createdStudent)
	}
}

func TestUpdateStudent(t *testing.T) {
	studentRepo := NewStudent(db)
	updateStudent := models.Student{
		FirstName: faker.Name(),
		LastName:  faker.Name(),
		Age:       15,
		Mail:      "new mail",
		Id:        "a0e27142-e55e-44ab-b292-5880f79b4243",
	}
	id, err := studentRepo.UpdateSt(updateStudent)
	if assert.NoError(t, err) {
		updatedStudent, err := studentRepo.GetById(id)
		if assert.NoError(t, err) {
			assert.Equal(t, updateStudent.FirstName, updatedStudent.FirstName)
			assert.Equal(t, updateStudent.LastName, updatedStudent.LastName)
			assert.Equal(t, updateStudent.Age, updatedStudent.Age)
			assert.Equal(t, updateStudent.Mail, updatedStudent.Mail)
		}
	}
}

func TestDeleteStudent(t *testing.T) {
	studentRepo := NewStudent(db)

	id := "a0e27142-e55e-44ab-b292-5880f79b4243"

	erDr := studentRepo.DeleteSt(id)
	if assert.NoError(t, erDr) {
		_, err := studentRepo.GetById(id)
		assert.Error(t, err, err)
	}
}

func TestGetAllStudent(t *testing.T) {
	studentRepo := NewStudent(db)
	reqStudent := models.Student{
		FirstName: faker.Name(),
		Age:       10,
		LastName:  faker.Word(),
	}

	getAllreq := models.GetAllStudentsRequest{
		Search: "",
		Page:   1,
		Limit:  10,
	}

	_, err := studentRepo.Create(reqStudent)
	oldCount, _ := studentRepo.GetAll(getAllreq)

	if assert.NoError(t, err) {
		_, err = studentRepo.Create(reqStudent)
		if assert.NoError(t, err) {
			newcount, _ := studentRepo.GetAll(getAllreq)
			assert.Equal(t, 1, newcount.Count-oldCount.Count)
		}
	}
}
