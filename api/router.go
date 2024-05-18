package api

import (
	"backend_course/lms/api/handler"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/service"
	// "net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// New ...
// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(service service.IServiceManager, log logger.ILogger) *gin.Engine {
	h := handler.NewStrg(service, log)

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// r.Use(authMiddleware)

	r.POST("/student", h.CreateStudent)
	r.GET("/student", h.GetAllStudents)
	r.GET("/student/:id", h.GetById)
	r.PUT("/student/updatestudent/:id", h.UpdateStudent)
	r.DELETE("/student/deletstudent/:id", h.DeleteStudent)
	r.GET("/student/status/:id", h.StudentStatus)
	r.GET("student/lesson/:id",h.StudentLesson)

	r.POST("/teacher", h.CreateTeacher)
	r.PUT("/teacher/updateteacher/:id", h.UpdateTeacher)
	r.GET("/teacher", h.GetAllTeacher)
	r.GET("/teacher/:id", h.GetByIdTeacher)
	r.DELETE("/teacher/deleteteacher/:id", h.DeleteTeacher)
	r.GET("/teacher/lesson/:id",h.TeacherLesson)
	r.POST("/teacher/login",h.TeacherLogin)
	r.POST("/teacher/register",h.TeacherRegister)
	r.POST("/teacher/register-confirm",h.TeacherRegisterConfirm)
	r.POST("/teacher/loginbymail",h.TeacherLoginByMail)
	r.POST("/teacher/login-confirm",h.TeacherLoginConfirm)

	r.POST("/subject", h.CreateSubject)
	r.GET("/subject/:id", h.GetbyIdSubject)
	r.PUT("/subject/updatesubject/:id", h.UpdateSubject)
	r.DELETE("/subject/deletsubject/:id", h.DeleteSubject)
	r.GET("/subject", h.GetAllSubject)

	r.POST("/timetable", h.CreateTimeTable)
	r.DELETE("/timetable/:id", h.DeleteTimeTable)
	r.POST("/timetable/studentsattandence",h.GetAllStudentsAttandenceReport)

	return r
}

// func authMiddleware(c *gin.Context) {
// 	if c.GetHeader("Authorization") != "secret" {
// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
// 		return
// 	}
// 	c.Next()
// }