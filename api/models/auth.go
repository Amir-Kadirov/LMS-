package models

type LoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginRequestEmail struct {
	Login string `json:"login"`
	Otp string `json:"otp"`
}

type AuthInfo struct {
	UserID   string `json:"user_id"`
	UserRole string `json:"user_role"`
}

type RegisterRequest struct {
	Mail string `json:"mail"`
}

type TeacherRegister struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	SubjectId string `json:"subject_id"`
	StartWork string `json:"start_work"`
	Mail      string `json:"mail"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type TeacherRegisterConfirm struct {
	Teacher Teacher `json:"teacher"`
	Otp     string  `json:"otp"`
}
