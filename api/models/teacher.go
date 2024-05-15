package models

type Teacher struct {
	Id        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	SubjectId string `json:"subject_id"`
	StartWork string `json:"start_work"`
	Mail      string `json:"mail"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type GetAllTeacherResponse struct {
	Teachers []Teacher `json:"teacher"`
	Count    int       `json:"count"`
}

type TeacherStudent struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TeacherSubjects struct {
	Type string `json:"type"`
	Name string `josn:"name"`
}

type GetByIdTeacher struct {
	Id        string             `json:"id"`
	FirstName string             `json:"first_name"`
	LastName  string             `json:"last_name"`
	SubjectId string             `json:"subject_id"`
	StartWork string             `json:"start_work"`
	Mail      string             `json:"mail"`
	Phone     string             `json:"phone"`
	Student   []TeacherStudent   `json:"student"`
	Subjects  []TeacherSubjects  `json:"subject"`
	TimeTable []TeacherTimeTable `json:"timetable"`
}

type TeacherTimeTable struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type SliceTeacherStudents struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}

type CheckLessonTeacher struct {
	StudentName string                 `json:"teacher_name"`
	StartDate   string                 `json:"start_date"`
	EndDate     string                 `json:"end_date"`
	SubjectName string                 `json:"subject_name"`
	Student     []SliceTeacherStudents `json:"students"`
}
