package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/storage"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
func New(store storage.IStorage) *gin.Engine {
	h := handler.NewStrg(store)

	r := gin.Default()

	r.POST("/student", h.CreateStudent)
	r.GET("/student", h.GetAllStudents)
	r.GET("/student/:id", h.GetById)
	r.PUT("/student/updatestudent/:id", h.UpdateStudent)
	r.PUT("/student/updatepassword/:id/:password", h.UpdateStudentPassword)
	r.DELETE("/student/deletstudent/:id", h.DeleteStudent)
	r.GET("/student/status/:id", h.StudentStatus)

	r.POST("/teacher", h.CreateTeacher)
	r.PUT("/teacher/updateteacher/:id", h.UpdateTeacher)
	r.GET("/teacher", h.GetAllTeacher)
	r.GET("/teacher/:id", h.GetByIdTeacher)
	r.DELETE("/teacher/deleteteacher/:id",h.DeleteTeacher)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}