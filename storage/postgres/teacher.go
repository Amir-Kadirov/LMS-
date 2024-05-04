package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type teacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacher(db *pgxpool.Pool) teacherRepo {
	return teacherRepo{
		db: db,
	}
}

func (t teacherRepo) CreateTeacher(teacher models.Teacher) (string, error) {
	id := uuid.New()
	subjectId := uuid.New()
	query := `INSERT INTO teacher(id,first_name,last_name,subject_id,start_work,phone,mail,created_at)
	 VALUES($1,$2,$3,$4,$5,$6,$7,NOW())`

	_, err := t.db.Exec(context.Background(), query, id, teacher.FirstName, teacher.LastName,
		subjectId, teacher.StartWork, teacher.Phone, teacher.Mail)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (t teacherRepo) UpdateTeacher(teacher models.Teacher) error {
	query := `UPDATE teacher SET first_name=$1,last_name=$2,start_work=$3,phone=$4,mail=$5,updated=NOW() WHERE id=$6`

	_, err := t.db.Exec(context.Background(), query, teacher.FirstName, teacher.LastName,
		teacher.StartWork, teacher.Phone, teacher.Mail, teacher.Id)
	if err != nil {
		return err
	}

	return nil
}

func (t teacherRepo) GetAllTeacher(req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error) {
	resp := models.GetAllTeacherResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					first_name,
					last_name,
					subject_id,
					start_work,
					phone,
					mail
				FROM teacher
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := t.db.Query(context.Background(), query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			teacher  models.Teacher
			lastName sql.NullString
		)
		if err := rows.Scan(
			&teacher.Id,
			&teacher.FirstName,
			&lastName,
			&teacher.SubjectId,
			&teacher.StartWork,
			&teacher.Phone,
			&teacher.Mail); err != nil {
			return resp, err
		}

		teacher.LastName = pkg.NullStringToString(lastName)
		resp.Teachers = append(resp.Teachers, teacher)
	}

	err = t.db.QueryRow(context.Background(), `SELECT count(*) from teacher WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t teacherRepo) GetTeacherbyId(id string) (models.Teacher, error) {
	resp := models.Teacher{}
	query := `SELECT id,
	first_name,
	last_name,
	subject_id,
	start_work,
	phone,
	mail
       FROM teacher WHERE id=$1`

	row := t.db.QueryRow(context.Background(), query, id)
	err := row.Scan(&resp.Id, &resp.FirstName, &resp.LastName, &resp.SubjectId, &resp.StartWork, &resp.Phone, &resp.Mail)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t teacherRepo) DeleteTeacher(id string) error {
	query := `DELETE FROM teacher WHERE id=$1`

	_, err := t.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}
