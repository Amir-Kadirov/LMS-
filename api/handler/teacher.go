package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateNumber(teacher.Phone); err != nil {
		handleResponse(c, "error while validating teacher phone: "+teacher.Phone, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.TeacherStorage().CreateTeacher(teacher)
	if err != nil {
		handleResponse(c, "error while creating teacher", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

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

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	teacher.Id = id

	err := h.Store.TeacherStorage().UpdateTeacher(teacher)
	if err != nil {
		handleResponse(c, fmt.Sprintf("error while updating teacher %s", teacher.Id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

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
		handleResponse(c, "error while parsing page", http.StatusBadRequest, err.Error())
		return
	}
	limit, err := ParseLimitQueryParam(c)
	if err != nil {
		handleResponse(c, "error while parsing limit", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.Store.TeacherStorage().GetAllTeacher(models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, "error while getting all teacher", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "Select all successful", http.StatusOK, resp)
}

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

	data, err := h.Store.TeacherStorage().GetTeacherbyId(id)
	if err != nil {
		handleResponse(c, fmt.Sprintf("error while get by id teacher %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Select by id successful", http.StatusOK, data)
}

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

	err := h.Store.TeacherStorage().DeleteTeacher(id)
	if err != nil {
		handleResponse(c, "error while deleting teacher", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Student deleted successfully", http.StatusOK, err)
}