package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/storage"
	"context"
	"fmt"
)

type timetableService struct {
	storage storage.IStorage
}

func NewTimeTableService(storage storage.IStorage) timetableService {
	return timetableService{storage: storage}
}

func (t timetableService) CreateTimeTable(ctx context.Context,timetable models.TimeTable) (string, error) {

	id, err := t.storage.TimeTableStorage().CreateTimeTable(ctx,timetable)
	if err != nil {
		fmt.Println("error while creating timetable, err: ", err)
		return "", err
	}

	return id, nil
}

func (t timetableService) DeleteTimeTable(ctx context.Context,id string) error {
	err := t.storage.TimeTableStorage().DeleteTimeTable(ctx,id)
	if err != nil {
		fmt.Println("error while deleting time table: ", err)
		return err
	}

	return nil
}
