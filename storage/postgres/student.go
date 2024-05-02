package postgres

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg"
	"database/sql"
	"math/rand"
	"strconv"

	"github.com/google/uuid"
)

type studentRepo struct {
	db *sql.DB
}

func NewStudent(db *sql.DB) studentRepo {
	return studentRepo{
		db: db,
	}
}

func (s *studentRepo) Create(student models.Student) (string, error) {

	id := uuid.New()
	student.ExternalId = strconv.Itoa(rand.Intn(999))

	query := ` INSERT INTO students (id, first_name,last_name,age,external_id,
		phone,mail,pasword, created_at) VALUES ($1, $2,$3,$4,$5,$6,$7,$8, NOW()) `

	_, err := s.db.Exec(query, id, student.FirstName, student.LastName, student.Age, student.ExternalId,
		student.Phone, student.Mail, student.Pasword)
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
	rows, err := s.db.Query(query, offest, req.Limit)
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

	err = s.db.QueryRow(`SELECT count(*) from students WHERE TRUE ` + filter + ``).Scan(&resp.Count)
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

	_, err := s.db.Exec(query,
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

func (s *studentRepo) UpdateStPassword(student models.Student) (string, error) {

	query := ` UPDATE students set pasword = $1,updated = NOW() WHERE id = $2 `

	_, err := s.db.Exec(query, student.Pasword, student.Id)
	if err != nil {
		return "", err
	}

	return student.Id, nil
}

func (s *studentRepo) GetById(student models.GetByIdRequest) (models.GetStudent, error) {
	resp := models.GetStudent{}

	query := `select id,first_name,last_name,age,external_id,phone,mail from students where external_id=$1`

	row := s.db.QueryRow(query, student.ExternalId)

	err := row.Scan(&resp.Id, &resp.FirstName, &resp.LastName, &resp.Age, &resp.ExternalId, &resp.Phone, &resp.Mail)

	if err != nil {
		return resp, err
	}

	return resp, nil
}


func (s *studentRepo) DeleteSt(student models.Student) (string, error) {

	query := `DELETE FROM students WHERE id = $1 `

	_, err := s.db.Exec(query,student.Id)
	if err != nil {
		return "", err
	}

	return student.Id, nil
}