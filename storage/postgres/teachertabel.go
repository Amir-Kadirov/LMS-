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

func (t timetableREpo) CreateTimeTable(ctx context.Context, timeTable models.TimeTable) (string, error) {
	id := uuid.New()

	query := `INSERT INTO time_tables(id,teacher_id,student_id,subject_id,start_date,end_date) VALUES($1,$2,$3,$4,$5,$6)`
	_, err := t.db.Exec(ctx, query, id, timeTable.TeacherId, timeTable.StudentId, timeTable.SubjectId, timeTable.FromDate, timeTable.ToDate)
	if err != nil {
		return "", err
	}

	return id.String(), nil
}

func (t timetableREpo) DeleteTimeTable(ctx context.Context, id string) error {
	query := `DELETE FROM time_tables WHERE id=$1`
	_, err := t.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *timetableREpo) GetAllStudentsAttandenceReport(ctx context.Context, req models.GetAllStudentsAttandenceReportRequest) (models.GetAllStudentsAttandenceReportResponse, error) {
	resp := models.GetAllStudentsAttandenceReportResponse{}
	filter := ""
	offest := (req.Page - 1) * req.Limit

	if req.StudentId != "" {
		filter = ` AND s.id =` + req.StudentId + ` `
	}

	if req.TeacherId != "" {
		filter += ` AND t.id =` + req.TeacherId + ` `
	}

	if req.StartDate != "" && req.EndDate != "" {
		filter += ` AND tt.start_date BETWEEN '` + req.StartDate + `' AND '` + req.EndDate + `' `
	}

	// 	1. Student name,
	// 	2. student createdAt,
	// 	3. o’qituvchi name,
	// 	4. studying_time,
	// 	5. avg_studying_time,
	query := `
	SELECT
                s.id,
                s.first_name || ' ' || s.last_name AS student_name,
                TO_CHAR(s.created_at,'YYYY-MM-DD HH:MM:SS'),
                t.first_name || ' ' || t.last_name AS teacher_name,
                EXTRACT(EPOCH FROM (tt.end_date - tt.start_date)) / 60 AS studying_time
                
        FROM 
                time_tables tt
                JOIN students s on tt.student_id = s.id
                JOIN teacher t on tt.teacher_id = t.id

	WHERE 
		TRUE ` + filter + `
	OFFSET $1 LIMIT $2
					`
	rows, err := s.db.Query(ctx, query, offest, req.Limit)
	if err != nil {
		return resp, err
	}
	studentAttandence := models.StudentAttandenceReport{}
	for rows.Next() {

		if err := rows.Scan(
			&studentAttandence.StudentId,
			&studentAttandence.StudentName,
			&studentAttandence.StudentCreatedAt,
			&studentAttandence.TeacherName,
			&studentAttandence.StudyTime); err != nil {
			return resp, err
		}

		resp.Students = append(resp.Students, studentAttandence)
	}

	err = s.db.QueryRow(ctx, `SELECT COUNT(*) from time_table time_table tt
	JOIN students s on tt.student_id = s.id
	JOIN teachers t on tt.teacher_id = t.id WHERE TRUE `+filter+``).Scan(&resp.Count)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
