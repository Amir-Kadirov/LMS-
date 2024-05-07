package postgres

import (
	"backend_course/lms/api/models"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateTimeTable(t *testing.T) {
	TimeTableRepo := NewTimeTable(db)

	reqTimeTable := models.TimeTable{
		FromDate: faker.Timestamp(),
		ToDate: faker.Timestamp(),
		TeacherId: "21d80a3d-6be4-4497-a5cf-a0d472b6d246",
		SubjectId: "37f19649-eb98-4924-9ddb-be778f3bd3b8",
		StudentId: "107cbb6a-00fa-49d8-b959-a8a71aab72c5",
	}

	_, err := TimeTableRepo.CreateTimeTable(reqTimeTable)
	assert.NoError(t, err,"Created Time Table") 
	}

func TestDeleteTimeTable(t *testing.T) {
	TimeTableRepo := NewTimeTable(db)

	id := "a0e27142-e55e-44ab-b292-5880f79b4243"

	err := TimeTableRepo.DeleteTimeTable(id)
	if assert.NoError(t, err) {
		return
	}
}