package postgres

import (
	"backend_course/lms/api/models"
	"fmt"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)

	reqTeacher := models.Teacher{
		FirstName: faker.Name(),
		Phone: "+998901111111",
		LastName:  faker.Word(),
	}

	id, err := teacherRepo.CreateTeacher(reqTeacher)
	if assert.NoError(t, err) {
		createdTeacher, err := teacherRepo.GetTeacherbyId(id)
		if assert.NoError(t, err) {
			assert.Equal(t, reqTeacher.FirstName, createdTeacher.FirstName)
			assert.Equal(t, reqTeacher.Phone, createdTeacher.Phone)
			assert.Equal(t, reqTeacher.LastName, createdTeacher.LastName)

		} else {
			return
		}
		fmt.Println("Created Teacher", createdTeacher)
	}
}

func TestUpdateTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)
	updateTeacher := models.Teacher{
		FirstName: faker.Name(),
		LastName:  faker.Name(),
		Mail:      "new mail",
		Id:        "3ed68793-cd2a-4c0a-9358-0eb2faa05cf5",
	}
	id, err := teacherRepo.UpdateTeacher(updateTeacher)
	if assert.NoError(t, err) {
		updatedTeacher, err := teacherRepo.GetTeacherbyId(id)
		if assert.NoError(t, err) {
			assert.Equal(t, updateTeacher.FirstName, updatedTeacher.FirstName)
			assert.Equal(t, updateTeacher.LastName, updatedTeacher.LastName)
			assert.Equal(t, updateTeacher.Mail, updatedTeacher.Mail)
		}
	}
}

func TestDeleteTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)

	id := "a0e27142-e55e-44ab-b292-5880f79b4243"

	erDr := teacherRepo.DeleteTeacher(id)
	if assert.NoError(t, erDr) {
		_, err := teacherRepo.GetTeacherbyId(id)
		assert.Error(t, err, err)
	}
}

func TestGetAllTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)
	reqTeacher := models.Teacher{
		FirstName: faker.Name(),
		LastName:  faker.Word(),
	}

	getAllreq := models.GetAllStudentsRequest{
		Search: "",
		Page:   1,
		Limit:  10,
	}

	_, err := teacherRepo.CreateTeacher(reqTeacher)
	oldCount, _ := teacherRepo.GetAllTeacher(getAllreq)

	if assert.NoError(t, err) {
		_, err = teacherRepo.CreateTeacher(reqTeacher)
		if assert.NoError(t, err) {
			newcount, _ := teacherRepo.GetAllTeacher(getAllreq)
			assert.Equal(t, 1, newcount.Count-oldCount.Count)
		}
	}
}
