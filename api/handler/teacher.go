package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Security ApiKeyAuth
// @Router		/teacher [post]
// @Summary		Creates a teacher
// @Description	This api creates a teacher and returns its id
// @Tags		Teacher
// @Accept		json
// @Produce		json
// @Param		teacher body models.Teacher true "teacher"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTeacher(c *gin.Context) {
	teacher := models.Teacher{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if !check.ValidatePhone(teacher.Phone) {
		handleResponse(c, h.Log, "error while validating teacher phone: "+teacher.Phone, http.StatusBadRequest, errors.New("wrong phone for country Uzb"))
		return
	}

	if !check.ValidateGmail(teacher.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	if !check.ValidatePassword(teacher.Password) {
		handleResponse(c, h.Log, "error while validating teacher password", http.StatusBadRequest, errors.New("unsecure password"))
		return
	}

	id, err := h.Service.Teacher().CreateTeacher(c.Request.Context(), teacher)
	if err != nil {
		handleResponse(c, h.Log, "error while creating teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/teacher/updateteacher/{id} [put]
// @Summary		Update a teacher
// @Description	This API updates a Teacher
// @Tags		Teacher
// @Accept		json
// @Produce		json
// @Param		id path string true "Teacher ID"
// @Param		teacher body models.Teacher true "Teacher object to update"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) UpdateTeacher(c *gin.Context) {

	id := c.Param("id")

	teacher := models.Teacher{}

	_, err := getAuthInfo(c)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusUnauthorized, err.Error())
		return
	}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if !check.ValidatePhone(teacher.Phone) {
		handleResponse(c, h.Log, "error while validating teacher phone: "+teacher.Phone, http.StatusBadRequest, errors.New("wrong phone for country Uzb"))
		return
	}

	if !check.ValidateGmail(teacher.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	if !check.ValidatePassword(teacher.Password) {
		handleResponse(c, h.Log, "error while validating teacher password", http.StatusBadRequest, errors.New("unsecure password"))
		return
	}

	teacher.Id = id

	err = h.Service.Teacher().UpdateTeacher(c.Request.Context(), teacher)
	if err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("error while updating teacher %s", teacher.Id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/teacher [get]
// @Summary		Get a teacher
// @Description	This API returns all teacher
// @Tags		Teacher
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllTeacher(c *gin.Context) {
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

	resp, err := h.Service.Teacher().GetAllTeacher(c.Request.Context(), models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all teacher", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "Select all successful", http.StatusOK, resp)
}

// @Security ApiKeyAuth
// @Router		/teacher/{id} [get]
// @Summary		Get by id a teacher
// @Description	This API get by id a teacher
// @Tags		Teacher
// @Produce		json
// @Param		id path string true "Teacher id"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) GetByIdTeacher(c *gin.Context) {
	id := c.Param("id")

	data, err := h.Service.Teacher().GetTeacherbyId(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("error while get by id teacher %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Select by id successful", http.StatusOK, data)
}

// @Security ApiKeyAuth
// @Router		/teacher/deleteteacher/{id} [delete]
// @Summary		Delete a teacher
// @Description	This API delete a teacher
// @Tags		Teacher
// @Produce		json
// @Param		id path string true "Teacher ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) DeleteTeacher(c *gin.Context) {

	id := c.Param("id")

	err := h.Service.Teacher().DeleteTeacher(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "teacher deleted successfully", http.StatusOK, err)
}

// @Security ApiKeyAuth
// @Router		/teacher/lesson/{id} [get]
// @Summary		Check the lesson of a teacher
// @Description	This API endpoint checks the lesson of a teacher by their ID.
// @Tags		Teacher
// @Param		id path string true "teacher ID"
// @Produce		json
// @Success		200 {object} models.Response "Status checked successfully"
// @Failure		400 {object} models.Response "Bad Request"
// @Failure		404 {object} models.Response "Not Found"
// @Failure		500 {object} models.Response "Internal Server Error"
func (h Handler) TeacherLesson(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("Error while validating ID %s", id), http.StatusBadRequest, err.Error())
		fmt.Println(id)
		return
	}

	lesson, err := h.Service.Teacher().LessonTeacher(c.Request.Context(),id)
	if err != nil {
		handleResponse(c, h.Log, "Error while checking teacher lesson", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Lesson checked successfully", http.StatusOK, lesson)
}
