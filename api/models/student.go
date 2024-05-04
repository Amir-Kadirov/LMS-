package models

type Student struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	Pasword    string `json:"pasword"`
	Active     bool   `json:"active"`
}

type GetStudent struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail,omitempty"`
}

type GetAllStudentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllStudentsResponse struct {
	Students []GetStudent `json:"students"`
	Count    int64        `json:"count"`
}

type IsActiveResponse struct {
	Active bool `json:"active"`
}