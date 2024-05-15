package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/hash"
	"context"
	"database/sql"
	"errors"

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

func (t teacherRepo) CreateTeacher(ctx context.Context, teacher models.Teacher) (string, error) {
	id := uuid.New()
	subjectId := uuid.New()

	hashedPassword, err := hash.HashPassword(teacher.Password)
	if err != nil {
		return "", err
	}

	query := `INSERT INTO teacher(id,first_name,last_name,subject_id,start_work,phone,mail,created_at,password)
	 VALUES($1,$2,$3,$4,$5,$6,$7,NOW(),$8)`

	_, err = t.db.Exec(ctx, query, id, teacher.FirstName, teacher.LastName,
		subjectId, teacher.StartWork, teacher.Phone, teacher.Mail, hashedPassword)

	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (t teacherRepo) UpdateTeacher(ctx context.Context, teacher models.Teacher) (string, error) {
	query := `UPDATE teacher SET first_name=$1,last_name=$2,start_work=$3,phone=$4,mail=$5,password=$6,updated=NOW() WHERE id=$7`

	hashingPassword, err := hash.HashPassword(teacher.Password)
	if err != nil {
		return "", errors.New("error while hashing")
	}

	_, err = t.db.Exec(ctx, query, teacher.FirstName, teacher.LastName,
		teacher.StartWork, teacher.Phone, teacher.Mail, hashingPassword, teacher.Id)
	if err != nil {
		return "", err
	}

	return teacher.Id, nil
}

func (t teacherRepo) GetAllTeacher(ctx context.Context, req models.GetAllStudentsRequest) (models.GetAllTeacherResponse, error) {
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
	rows, err := t.db.Query(ctx, query, offest, req.Limit)
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

	err = t.db.QueryRow(ctx, `SELECT count(*) from teacher WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t teacherRepo) GetTeacherbyId(ctx context.Context, id string) (models.GetByIdTeacher, error) {
	resp := models.GetByIdTeacher{}
	query := `SELECT id,
	first_name,
	last_name,
	subject_id,
	start_work,
	phone,
	mail
       FROM teacher WHERE id=$1`

	row := t.db.QueryRow(ctx, query, id)
	err := row.Scan(&resp.Id, &resp.FirstName, &resp.LastName, &resp.SubjectId, &resp.StartWork, &resp.Phone, &resp.Mail)
	if err != nil {
		return resp, err
	}

	querySubTimeTab := `SELECT sub.name,
	                         sub.type,
							 to_char(tt.start_date,'YYYY-MM-DD HH:MM:SS'),
							 to_char(tt.end_date,'YYYY-MM-DD HH:MM:SS'),
							 s.first_name,
							 s.last_name
							 FROM time_tables tt 
							   INNER JOIN students s ON s.id=tt.student_id
							   INNER JOIN subjects sub ON sub.id=tt.subject_id
							   WHERE tt.teacher_id=$1`

	rows, err := t.db.Query(ctx, querySubTimeTab, resp.Id)
	if err != nil {
		return resp, err
	}
	for rows.Next() {
		student := models.TeacherStudent{}
		subject := models.TeacherSubjects{}
		timetable := models.TeacherTimeTable{}
		if err := rows.Scan(
			&subject.Name,
			&subject.Type,
			&timetable.StartDate,
			&timetable.EndDate,
			&student.FirstName,
			&student.LastName,
		); err != nil {
			return resp, err
		}
		resp.Subjects = append(resp.Subjects, subject)
		resp.TimeTable = append(resp.TimeTable, timetable)
		resp.Student = append(resp.Student, student)
	}
	return resp, nil
}

func (t teacherRepo) DeleteTeacher(ctx context.Context, id string) error {
	query := `DELETE FROM teacher WHERE id=$1`

	_, err := t.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *teacherRepo) CheckLessonTeacher(ctx context.Context, id string) (models.CheckLessonTeacher, error) {
	query := `SELECT s.first_name,
	               s.age,
				   to_char(tt.start_date,'YYYY-MM-DD HH:MM:SS'),
			       to_char(tt.end_date,'YYYY-MM-DD HH:MM:SS'),
				   sub.name
				   FROM time_tables tt INNER JOIN students s on s.id=tt.student_id
				                       INNER JOIN subjects sub on sub.id=tt.subject_id
									   WHERE tt.teacher_id=$1`

	resp := models.CheckLessonTeacher{}

	row, err := s.db.Query(ctx, query, id)

	if err == sql.ErrNoRows {
		return resp, errors.New("teacher didn't have lesson")
	} else if err != nil {
		return resp, err
	}

	for row.Next() {
		student := models.SliceTeacherStudents{}
		if err = row.Scan(&student.Name,
			&student.Age,
			&resp.StartDate,
			&resp.EndDate,
			&resp.SubjectName); err != nil {
			return resp, err
		}
		resp.Student = append(resp.Student, student)
	}

	return resp, nil
}

func (t teacherRepo) GetTeacherbyLogin(ctx context.Context, login string) (models.Teacher, error) {
	resp := models.Teacher{}
	query := `SELECT 
	id,
	first_name,
	last_name,
	subject_id,
	start_work,
	phone,
	mail,
	password
       FROM teacher 
	WHERE mail=$1`

	row := t.db.QueryRow(ctx, query, login)
	err := row.Scan(&resp.Id,
		&resp.FirstName,
		&resp.LastName,
		&resp.SubjectId,
		&resp.StartWork,
		&resp.Phone,
		&resp.Mail,
		&resp.Password)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
