package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

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

func (t subjectRepo) UpadateSubject(subject models.Subjects) (string, error) {
	query := `UPDATE subjects SET name=$1,type=$2,updated_at=NOW() where id=$3`

	_, err := t.db.Exec(context.Background(), query, subject.Name, subject.Type, subject.Id)
	if err != nil {
		return "", err
	}

	return subject.Id, nil
}

func (t subjectRepo) GetbyIdSubject(id string) (models.Subjects, error) {
	subjects := models.Subjects{}
	query := `SELECT id,name,type,to_char(created_at,'YYYY-MM-DD HH:MM:SS'),to_char(updated_at,'YYYY-MM-DD HH:MM:SS') FROM subjects where id=$1 `

	row := t.db.QueryRow(context.Background(), query, id)

	var (
		created_at sql.NullString
		updated_at sql.NullString
	)

	err := row.Scan(&subjects.Id, &subjects.Name, &subjects.Type, &created_at, &updated_at)
	if err != nil {
		return subjects, err
	}

	subjects.CreatedAt = pkg.NullStringToString(created_at)
	subjects.UpdatedAt = pkg.NullStringToString(updated_at)

	return subjects, nil
}

func (t subjectRepo) DeleteSubject(id string) error {
	query := `DELETE FROM subjects WHERE id=$1`
	_, err := t.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}
	return nil
}

func (t subjectRepo) GetAllSubject(req models.GetAllStudentsRequest) (models.SubjectGetAll, error) {
	resp := models.SubjectGetAll{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					name,
					type
				FROM subjects
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := t.db.Query(context.Background(), query, offest, req.Limit)

	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			Subject models.Subjects
			// lastName sql.NullString
		)
		if err := rows.Scan(
			&Subject.Id,
			&Subject.Name,
			&Subject.Type,
		); err != nil {
			return resp, err
		}

		// Subject.LastName = pkg.NullStringToString(lastName)
		resp.Subject=append(resp.Subject,Subject)
	}

	err = t.db.QueryRow(context.Background(), `SELECT count(*) from Subject WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
