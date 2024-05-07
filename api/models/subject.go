package models

type Subjects struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type SubjectGetAll struct{
	Subject []Subjects
	Count int
}