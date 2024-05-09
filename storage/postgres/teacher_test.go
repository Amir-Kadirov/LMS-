package postgres

import (
	"backend_course/lms/api/models"
	"context"
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

	id, err := teacherRepo.CreateTeacher(context.Background() ,reqTeacher)
	if assert.NoError(t, err) {
		createdTeacher, err := teacherRepo.GetTeacherbyId(context.Background(),id)
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
		Id:        "c63fce30-58e0-477a-aec9-47f9d2d6ef4b",
	}
	id, err := teacherRepo.UpdateTeacher(context.Background(),updateTeacher)
	if assert.NoError(t, err) {
		updatedTeacher, err := teacherRepo.GetTeacherbyId(context.Background(),id)
		if assert.NoError(t, err) {
			assert.Equal(t, updateTeacher.FirstName, updatedTeacher.FirstName)
			assert.Equal(t, updateTeacher.LastName, updatedTeacher.LastName)
			assert.Equal(t, updateTeacher.Mail, updatedTeacher.Mail)
		}
	}
}

func TestDeleteTeacher(t *testing.T) {
	teacherRepo := NewTeacher(db)

	id := "fc0fe842-09dd-4e9b-b105-8d740174bbbe"

	erDr := teacherRepo.DeleteTeacher(context.Background(),id)
	if assert.NoError(t, erDr) {
		_, err := teacherRepo.GetTeacherbyId(context.Background(),id)
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

	_, err := teacherRepo.CreateTeacher(context.Background(),reqTeacher)
	oldCount, _ := teacherRepo.GetAllTeacher(context.Background(),getAllreq)

	if assert.NoError(t, err) {
		_, err = teacherRepo.CreateTeacher(context.Background(),reqTeacher)
		if assert.NoError(t, err) {
			newcount, _ := teacherRepo.GetAllTeacher(context.Background(),getAllreq)
			assert.Equal(t, 1, newcount.Count-oldCount.Count)
		}
	}
}
