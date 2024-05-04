package models

type Teacher struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	SubjectId string `json:"subject_id"`
	StartWork string `json:"start_work"`
	Mail      string `json:"mail"`
	Phone     string `json:"phone"`
}

type GetAllTeacherResponse struct {
	Teachers []Teacher `json:"teacher"`
	Count    int       `json:"count"`
}