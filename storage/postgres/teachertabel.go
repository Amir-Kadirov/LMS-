package postgres

import (
	"backend_course/lms/api/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type timetableREpo struct {
	db *pgxpool.Pool
}

func NewTimeTable(db *pgxpool.Pool) timetableREpo {
	return timetableREpo{
		db: db,
	}
}

func (t timetableREpo) CreateTimeTable(timeTable models.TimeTable) (string, error) {
	id := uuid.New().String()

	query:=`INSERT INTO time_tables(id,teacher_id,student_id,subject_id,start_date,end_date) VALUES($1,$2,$3,$4,$5,$6)`
	_,err:=t.db.Exec(context.Background(),query,id,timeTable.TeacherId,timeTable.StudentId,timeTable.SubjectId,timeTable.FromDate,timeTable.ToDate)
	if err!=nil {
		return "",err
	}

	return id,nil
}

func (t timetableREpo) DeleteTimeTable(id string) error {
	query:=`DELETE FROM time_tables WHERE id=$1`
	_,err:=t.db.Exec(context.Background(),query,id)
	if err!=nil {
		return err
	}
	return nil
}