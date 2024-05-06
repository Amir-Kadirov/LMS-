package models

type TimeTable struct {
	Id        string `json:"id"`
	TeacherId string `jsom:"teacher_id"`
	StudentId string `json:"student_id"`
	SubjectId string `json:"subject_id"`
	FromDate  string `json:"from_date"`
	ToDate    string `json:"to_date"`
}
