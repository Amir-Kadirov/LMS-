package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// @Router		/student [post]
// @Summary		Creates a student
// @Description	This api creates a student and returns its id
// @Tags		Student
// @Accept		json
// @Produce		json
// @Param		student body models.Student true "student"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateStudent(c *gin.Context) {
	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateAge(student.Age); err != nil {
		handleResponse(c, h.Log, "error while validating student age, year: "+strconv.Itoa(student.Age), http.StatusBadRequest, err.Error())
		return
	}

	if !check.ValidateGmail(student.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	if !check.ValidatePhone(student.Phone) {
		handleResponse(c, h.Log, "error while validating student phone: "+student.Phone, http.StatusBadRequest, errors.New("wrong phone for country Uzb"))
		return
	}

	if !check.ValidatePassword(student.Pasword) {
		handleResponse(c, h.Log, "error while validating student password", http.StatusBadRequest, errors.New("unsecure password"))
		return
	}

	id, err := h.Service.Student().Create(c.Request.Context(), student)
	if err != nil {
		handleResponse(c, h.Log, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/student/updatestudent/{id} [PUT]
// @Summary		Update a student
// @Description	This API updates a student
// @Tags		Student
// @Accept		json
// @Produce		json
// @Param		id path string true "Student ID"
// @Param		student body models.Student true "Student object to update"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) UpdateStudent(c *gin.Context) {

	id := c.Param("id")

	student := models.Student{}

	if err := c.ShouldBindJSON(&student); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if !check.ValidateGmail(student.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	if !check.ValidatePhone(student.Phone) {
		handleResponse(c, h.Log, "error while validating teacher phone: "+student.Phone, http.StatusBadRequest, errors.New("wrong phone for country Uzb"))
		return
	}

	if !check.ValidatePassword(student.Pasword) {
		handleResponse(c, h.Log, "error while validating student password", http.StatusBadRequest, errors.New("unsecure password"))
		return
	}

	student.Id = id

	id, err := h.Service.Student().UpdateStudent(c.Request.Context(), student)
	if err != nil {
		handleResponse(c, h.Log, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/student [get]
// @Summary		Get a student
// @Description	This API returns all students
// @Tags		Student
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllStudents(c *gin.Context) {
	search := c.Query("search")

	page, err := ParsePageQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, h.Log, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Service.Student().GetAllStudent(c.Request.Context(), models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all students", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "Select all successful", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router		/student/{id} [get]
// @Summary		Get by id a student
// @Description	This API get by id a student
// @Tags		Student
// @Produce		json
// @Param		id path string true "Student Id"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.Service.Student().GetByIdStudent(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("error while get by id student %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Select by id successful", http.StatusOK, data)
}

// @Security ApiKeyAuth
// @Router		/student/deletstudent/{id} [delete]
// @Summary		Delete a student
// @Description	This API delete a student
// @Tags		Student
// @Produce		json
// @Param		id path string true "Student ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) DeleteStudent(c *gin.Context) {

	Id := c.Param("id")

	err := h.Service.Student().DeleteStudent(c.Request.Context(), Id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting  student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Student deleted successfully", http.StatusOK, err)
}

// @Security ApiKeyAuth
// @Router		/student/status/{id} [get]
// @Summary		Check the status of a student
// @Description	This API endpoint checks the status of a student by their ID.
// @Tags		Student
// @Param			id path string true "Student ID"
// @Produce		json
// @Success		200 {object} models.Response "Status checked successfully"
// @Failure		400 {object} models.Response "Bad Request"
// @Failure		404 {object} models.Response "Not Found"
// @Failure		500 {object} models.Response "Internal Server Error"
func (h Handler) StudentStatus(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("Error while validating ID %s", id), http.StatusBadRequest, err.Error())
		fmt.Println(id)
		return
	}

	boolean, err := h.Service.Student().StatusStudent(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "Error while checking student status", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Status checked successfully", http.StatusOK, boolean)
}

// @Security ApiKeyAuth
// @Router		/student/lesson/{id} [get]
// @Summary		Check the lesson of a student
// @Description	This API endpoint checks the lesson of a student by their ID.
// @Tags		Student
// @Param		id path string true "Student ID"
// @Produce		json
// @Success		200 {object} models.Response "Status checked successfully"
// @Failure		400 {object} models.Response "Bad Request"
// @Failure		404 {object} models.Response "Not Found"
// @Failure		500 {object} models.Response "Internal Server Error"
func (h Handler) StudentLesson(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("Error while validating ID %s", id), http.StatusBadRequest, err.Error())
		fmt.Println(id)
		return
	}

	lesson, err := h.Service.Student().LessonStudent(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "Error while checking student lesson", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Lesson checked successfully", http.StatusOK, lesson)
}