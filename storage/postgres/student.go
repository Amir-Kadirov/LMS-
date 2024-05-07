package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type studentRepo struct {
	db *pgxpool.Pool
}

func NewStudent(db *pgxpool.Pool) studentRepo {
	return studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(student models.Student) (string, error) {

	id := uuid.New()
	student.ExternalId = strconv.Itoa(rand.Intn(999))

	query := ` INSERT INTO students (id, first_name,last_name,age,external_id,
		phone,mail,pasword,active, created_at) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9, NOW()) `

	_, err := s.db.Exec(context.Background(), query, id, student.FirstName, student.LastName, student.Age, student.ExternalId,
		student.Phone, student.Mail, student.Pasword, student.Active)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (s *studentRepo) GetAll(req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
	resp := models.GetAllStudentsResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.Search != "" {
		filter = ` AND first_name ILIKE '%` + req.Search + `%' `
	}

	query := `SELECT id,
					first_name,
					last_name,
					age,
					external_id,
					phone,
					mail
				FROM students
				WHERE TRUE ` + filter + `
				OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(context.Background(), query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		var (
			student  models.GetStudent
			lastName sql.NullString
		)
		if err := rows.Scan(
			&student.Id,
			&student.FirstName,
			&lastName,
			&student.Age,
			&student.ExternalId,
			&student.Phone,
			&student.Mail); err != nil {
			return resp, err
		}

		student.LastName = pkg.NullStringToString(lastName)
		resp.Students = append(resp.Students, student)
	}

	err = s.db.QueryRow(context.Background(), `SELECT count(*) from students WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *studentRepo) UpdateSt(student models.Student) (string, error) {

	query := ` UPDATE students set first_name=$1,
								   last_name=$2,
								   age=$3,
								   phone=$4,
								   mail=$5,
								   updated = NOW()
								       WHERE id = $6 `

	_, err := s.db.Exec(context.Background(), query,
		student.FirstName,
		student.LastName,
		student.Age,
		student.Phone,
		student.Mail,
		student.Id)
	if err != nil {
		return "", err
	}

	return student.Id, nil
}

func (s *studentRepo) UpdateStPassword(id string, password string) (string, error) {

	query := ` UPDATE students set pasword = $1,updated = NOW() WHERE id = $2 `

	_, err := s.db.Exec(context.Background(), query, password, id)
	if err != nil {
		return "", err
	}

	return password, nil
}

func (s *studentRepo) GetById(id string) (models.GetStudent, error) {
	resp := models.GetStudent{}

	query := `SELECT id,
	          first_name,
			  last_name,
			  age,
			  external_id,
			  phone,
			  mail 
			  FROM students WHERE id=$1`

	row := s.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&resp.Id, &resp.FirstName, &resp.LastName, &resp.Age, &resp.ExternalId, &resp.Phone, &resp.Mail)

	if err != nil {
		fmt.Println("error")
		return resp, err
	}

	querySubTimeTab := `SELECT sub.name,
	                         sub.type,
							 to_char(tt.start_date,'YYYY-MM-DD HH:MM:SS'),
							 to_char(tt.end_date,'YYYY-MM-DD HH:MM:SS'),
							 t.first_name 
							 FROM time_tables tt 
							   INNER JOIN teacher t ON t.id=tt.teacher_id
							   INNER JOIN subjects sub ON sub.id=tt.subject_id
							   WHERE tt.student_id=$1`

	rows, err := s.db.Query(context.Background(), querySubTimeTab, resp.Id)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		teacher := models.StudentTeacher{}
		subject := models.StudentSubjects{}
		timetable := models.StudentTimeTable{}
		if err := rows.Scan(
			&subject.Name,
			&subject.Type,
			&timetable.StartDate,
			&timetable.EndDate,
			&teacher.Name,
		); err != nil {
			return resp, err
		}
		resp.Subjects = append(resp.Subjects, subject)
		resp.TimeTable = append(resp.TimeTable, timetable)
		resp.Teacher = append(resp.Teacher, teacher)
	}
	return resp,nil
}

func (s *studentRepo) DeleteSt(id string) error {

	query := `DELETE FROM students WHERE id = $1 `

	_, err := s.db.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepo) StatusSt(id string) (models.IsActiveResponse, error) {

	req := models.IsActiveResponse{}

	query := `SELECT active FROM students where id=$1`

	row := s.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&req.Active)
	if err != nil {
		return req, err
	}

	return req, nil
}
