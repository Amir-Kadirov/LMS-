package postgres

import (
	"backend_course/lms/api/models"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type subjectRepo struct {
	db *pgxpool.Pool
}

func NewSubject(db *pgxpool.Pool) subjectRepo {
	return subjectRepo{
		db: db,
	}
}

func (t subjectRepo) CreateSubject(subject models.Subjects) (string, error) {
	id := uuid.New()

	query := `INSERT INTO subjects(id,name,type,created_at) VALUES($1,$2,$3,NOW())`

	_, err := t.db.Exec(context.Background(), query, id, subject.Name, subject.Type)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (t subjectRepo) UpadateSubject(subject models.Subjects) error {
	query := `UPDATE subjects SET name=$1,type=$2,updated_at=NOW() where id=$3`

	_, err := t.db.Exec(context.Background(), query, subject.Name, subject.Type, subject.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t subjectRepo) GetbyIdSubject(id string) (models.Subjects, error) {
	subjects:=models.Subjects{}
	query := `SELECT id,name,type,created_at,updated_at FROM subjects where id=$1 `

	row:=t.db.QueryRow(context.Background(),query,id)
	
	err:=row.Scan(&subjects.Id,&subjects.Name,&subjects.Type,&subjects.CreatedAt,subjects.UpdatedAt)
	if err!=nil {
		return subjects,err
	}

	return subjects,nil

}

func (t subjectRepo) DeleteSubject(id string) error {
	query:=`DELETE FROM subjects WHERE id=$1`
	_,err:=t.db.Exec(context.Background(),query,id)
	if err!=nil {
		return err
	}
	return nil
}