package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"
)

func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/student", h.CreateStudent)
	r.GET("/student", h.GetAllStudents)
	r.GET("/student/:id", h.GetById)
	r.PUT("/student/updatestudent/:id", h.UpdateStudent)
	r.PUT("/student/updatepassword/:id", h.UpdateStudentPassword)
	r.DELETE("/student/deletstudent/:id", h.DeleteStudent)
	return r
}