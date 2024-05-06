package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"fmt"
)

type timetableService struct {
	storage storage.IStorage
}

func NewTimeTableService(storage storage.IStorage) timetableService {
	return timetableService{storage: storage}
}

func (t timetableService) CreateTimeTable(timetable models.TimeTable) (string, error) {

	id, err := t.storage.TimeTableStorage().CreateTimeTable(timetable)
	if err != nil {
		fmt.Println("error while creating timetable, err: ", err)
		return "", err
	}

	return id, nil
}

func (t timetableService) DeleteTimeTable(id string) error {
	err := t.storage.TimeTableStorage().DeleteTimeTable(id)
	if err != nil {
		fmt.Println("error while deleting time table: ", err)
		return err
	}

	return nil
}
