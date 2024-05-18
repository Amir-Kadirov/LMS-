package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/hash"
	smtp "backend_course/lms/pkg/helper"
	"context"
	"database/sql"
	"errors"
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

func (s *studentRepo) Create(ctx context.Context, student models.Student) (string, error) {

	id := uuid.New()
	student.ExternalId = strconv.Itoa(rand.Intn(999))

	hashedPassword, err := hash.HashPassword(student.Pasword)
	if err != nil {
		fmt.Println("in hash")
		return "", err
	}

	query := ` INSERT INTO students (id, first_name,last_name,age,external_id,
		phone,mail,pasword,active, created_at) VALUES ($1, $2,$3,$4,$5,$6,$7,$8,$9, NOW())`

	_, err = s.db.Exec(ctx, query, id, student.FirstName, student.LastName, student.Age, student.ExternalId,
		student.Phone, student.Mail, hashedPassword, student.Active)
	if err != nil {
		fmt.Println("in student insert")
		return "", err
	}
	
	for _, phone := range student.Phones {
        queryPhones := `INSERT INTO student_numbers (phone, student_id, created_at) VALUES ($1, $2, NOW())`
        _, err = s.db.Exec(ctx, queryPhones, phone.Phone, id)
        if err != nil {
            fmt.Println("in student number insert")
            return "", err
        }
    }

	if err!=nil {
		fmt.Println("in student number insert")
		return "",err
	}

	err = smtp.SendMail(student.Mail, "Welcome new student")
	if err != nil {
		fmt.Println("in send main")
		return "", err
	}

	return id.String(), nil
}

func (s *studentRepo) GetAll(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllStudentsResponse, error) {
    resp := models.GetAllStudentsResponse{}
    filter := ""
    offset := (req.Page - 1) * req.Limit

    if req.Search != "" {
        filter = ` AND first_name ILIKE '%` + req.Search + `%' `
    }

    query := `
        SELECT s.id, s.first_name, s.last_name, s.age, s.external_id, s.phone, s.mail, sn.phone
        FROM students s
        JOIN student_numbers sn ON s.id = sn.student_id
        WHERE TRUE ` + filter + `
        OFFSET $1 LIMIT $2
    `

    rows, err := s.db.Query(ctx, query, offset, req.Limit)
    if err != nil {
        return resp, err
    }
    defer rows.Close()

    for rows.Next() {
        var student models.GetStudent
        var lastName sql.NullString
        var phone models.StudentNumbers

        if err := rows.Scan(
            &student.Id,
            &student.FirstName,
            &lastName,
            &student.Age,
            &student.ExternalId,
            &student.Phone,
            &student.Mail,
            &phone.Phone,
        ); err != nil {
            return resp, err
        }

        student.LastName = pkg.NullStringToString(lastName)
        student.Phones = append(student.Phones, phone)
        resp.Students = append(resp.Students, student)
    }

    if err := rows.Err(); err != nil {
        return resp, err
    }

	err = s.db.QueryRow(ctx, `SELECT count(*) from students WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (s *studentRepo) UpdateSt(ctx context.Context, student models.Student) (string, error) {

	query := ` UPDATE students set first_name=$1,
								   last_name=$2,
								   age=$3,
								   phone=$4,
								   mail=$5,
								   pasword=$6,
								   updated = NOW()
								       WHERE id = $7 `

	_, err := s.db.Exec(ctx, query,
		student.FirstName,
		student.LastName,
		student.Age,
		student.Phone,
		student.Mail,
		student.Pasword,
		student.Id)
	if err != nil {
		return "", err
	}

	return student.Id, nil
}

func (s *studentRepo) GetById(ctx context.Context, id string) (models.GetStudent, error) {
	resp := models.GetStudent{}

	query := `SELECT s.id,
	          s.first_name,
			  s.last_name,
			  s.age,
			  s.external_id,
			  s.phone,
			  s.mail,
			  sn.phone
			  FROM students s
			  INNER JOIN student_numbers sn on sn.student_id=s.id
			  WHERE s.id=$1`

	row,err := s.db.Query(ctx, query, id)
	if err!=nil {
		return resp,err
	}

	for row.Next(){
		phone:=models.StudentNumbers{}
		if err := row.Scan(&resp.Id,
			&resp.FirstName,
			&resp.LastName,
			&resp.Age, 
			&resp.ExternalId, 
			&resp.Phone, 
			&resp.Mail,
			&phone.Phone,
			)
			err !=nil{
				return resp,err
			}
			resp.Phones = append(resp.Phones, phone)
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

	rows, err := s.db.Query(ctx, querySubTimeTab, resp.Id)
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
	return resp, nil
}

func (s *studentRepo) DeleteSt(ctx context.Context, id string) error {

	query := `DELETE FROM students WHERE id = $1 `

	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *studentRepo) StatusSt(ctx context.Context, id string) (models.IsActiveResponse, error) {

	req := models.IsActiveResponse{}

	query := `SELECT active FROM students where id=$1`

	row := s.db.QueryRow(ctx, query, id)

	err := row.Scan(&req.Active)
	if err != nil {
		return req, err
	}

	return req, nil
}

func (s *studentRepo) CheckLessonStudent(ctx context.Context, id string) (models.CheckLessonStudent, error) {
	query := `SELECT t.first_name,
				   to_char(tt.start_date,'YYYY-MM-DD HH:MM:SS'),
			       to_char(tt.end_date,'YYYY-MM-DD HH:MM:SS'),
				   sub.name
				   FROM time_tables tt INNER JOIN teacher t on t.id=tt.teacher_id
				                       INNER JOIN subjects sub on sub.id=tt.subject_id
									   WHERE tt.student_id=$1`

	lessons := models.CheckLessonStudent{}

	row := s.db.QueryRow(ctx, query, id)
	err := row.Scan(&lessons.TeacherName, &lessons.StartDate, &lessons.EndDate, &lessons.SubjectName)
	if err == sql.ErrNoRows {
		return lessons, errors.New("student didn't have lesson")
	} else if err != nil {
		return lessons, err
	}

	return lessons, nil
}