package postgres

import (
	"backend_course/lms/api/models"
	"context"
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

	id, err := studentRepo.Create(context.Background(),reqStudent)
	if assert.NoError(t, err) {
		createdStudent, err := studentRepo.GetById(context.Background(),id)
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
		Id:        "7dfc00b2-376d-4db8-80ad-eb35e194c561",
	}

	id, err := studentRepo.UpdateSt(context.Background(),updateStudent)
	if assert.NoError(t, err) {
		updatedStudent, err := studentRepo.GetById(context.Background(),id)
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

	id := "68afa24d-959c-44ce-add7-6e5974e04b37"

	erDr := studentRepo.DeleteSt(context.Background(),id)
	if assert.NoError(t, erDr) {
		_, err := studentRepo.GetById(context.Background(),id)
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

	_, err := studentRepo.Create(context.Background(),reqStudent)
	oldCount, _ := studentRepo.GetAll(context.Background(),getAllreq)

	if assert.NoError(t, err) {
		_, err = studentRepo.Create(context.Background(),reqStudent)
		if assert.NoError(t, err) {
			newcount, _ := studentRepo.GetAll(context.Background(),getAllreq)
			assert.Equal(t, 1, newcount.Count-oldCount.Count)
		}
	}
}

func TestStatusStudent(t *testing.T) {
    stRepo := NewStudent(db)
    isactive, err := stRepo.StatusSt(context.Background(), "7dfc00b2-376d-4db8-80ad-eb35e194c561")
    if assert.NoError(t, err) {
        if isactive.Active {
            assert.True(t, isactive.Active, "Expected student to be active")
        } else {
            assert.False(t, isactive.Active, "Expected student to be inactive")
        }
    }
}


func TestUpdatePassword(t *testing.T) {
	studentRepo := NewStudent(db)

	id := "68afa24d-959c-44ce-add7-6e5974e04b37"
	password:="New Password"

	newpassword,erDr := studentRepo.UpdateStPassword(context.Background(),id,password)
	if assert.NoError(t, erDr) {
		assert.Equal(t,password,newpassword)
	}
}