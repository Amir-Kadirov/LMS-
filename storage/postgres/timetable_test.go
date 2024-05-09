package postgres

import (
	"backend_course/lms/api/models"
	"context"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimeTable(t *testing.T) {
	TimeTableRepo := NewTimeTable(db)

	reqTimeTable := models.TimeTable{
		FromDate: faker.Timestamp(),
		ToDate: faker.Timestamp(),
		TeacherId: "90c97fc8-6c37-471f-b52a-4636d8036a4c",
		SubjectId: "4cae2595-12d8-424e-81bb-0bd7507ddde1",
		StudentId: "4a481ea4-777d-406d-8882-969716570ca1",
	}

	_, err := TimeTableRepo.CreateTimeTable(context.Background(),reqTimeTable)
	assert.NoError(t, err,"Created Time Table") 
	}

func TestDeleteTimeTable(t *testing.T) {
	TimeTableRepo := NewTimeTable(db)

	id := "64dc0907-c7cf-4b98-8480-9cd23edd4d20"

	err := TimeTableRepo.DeleteTimeTable(context.Background(),id)
	if assert.NoError(t, err) {
		return
	}
}