package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Router		/student [post]
// @Summary		Creates a student
// @Description	This api creates a student and returns its id
// @Tags		student
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
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}
	if err := check.ValidateAge(student.Age); err != nil {
		handleResponse(c, "error while validating student age, year: "+strconv.Itoa(student.Age), http.StatusBadRequest, err.Error())
		return
	}

	if err := check.ValidateNumber(student.Phone); err != nil {
		handleResponse(c, "error while validating student phone: "+student.Phone, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.StudentStorage().Create(student)
	if err != nil {
		handleResponse(c, "error while creating student", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

// @Router		/student/updatestudent/{id} [put]
// @Summary		Update a student
// @Description	This API updates a student
// @Tags		student
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
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	student.Id = id

	id, err := h.Store.StudentStorage().UpdateSt(student)
	if err != nil {
		handleResponse(c, "error while updating student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

// @Router		/student [get]
// @Summary		Get a student
// @Description	This API returns all students
// @Tags		student
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllStudents(c *gin.Context) {
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

	resp, err := h.Store.StudentStorage().GetAll(models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, "error while getting all students", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, "Select all successful", http.StatusOK, resp)
}

// @Router		/student/updatepassword/{id}/{password} [put]
// @Summary		Update a student's password
// @Description	This API endpoint updates a student's password by their ID.
// @Tags			student
// @Produce		json
// @Param			id path string true "Student ID"
// @Param			password path string true "New Password"
// @Success		200 {object} models.Response "Password updated successfully"
// @Failure		400 {object} models.Response "Bad Request"
// @Failure		404 {object} models.Response "Not Found"
// @Failure		500 {object} models.Response "Internal Server Error"
func (h Handler) UpdateStudentPassword(c *gin.Context) {
	id := c.Param("id")
	password := c.Param("password")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, "Error while validating student ID", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Store.StudentStorage().UpdateStPassword(id, password)
	if err != nil {
		handleResponse(c, "Error while updating student password", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Password updated successfully", http.StatusOK, id)
}

// @Router		/student/{id} [get]
// @Summary		Get by id a student
// @Description	This API get by id a student
// @Tags		student
// @Produce		json
// @Param		id path string true "Student ExternalId"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) GetById(c *gin.Context) {
	id := c.Param("id")

	data, err := h.Store.StudentStorage().GetById(id)
	if err != nil {
		handleResponse(c, fmt.Sprintf("error while get by id student %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Select by id successful", http.StatusOK, data)
}

// @Router		/student/deletstudent/{id} [delete]
// @Summary		Delete a student
// @Description	This API delete a student
// @Tags		student
// @Produce		json
// @Param		id path string true "Student ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) DeleteStudent(c *gin.Context) {

	ExternalId:= c.Param("id")

	err := h.Store.StudentStorage().DeleteSt(ExternalId)
	if err != nil {
		handleResponse(c, "error while deleting  student", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Student deleted successfully", http.StatusOK, err)
}

// @Router		/student/status/{id} [get]
// @Summary		Check the status of a student
// @Description	This API endpoint checks the status of a student by their ID.
// @Tags			student
// @Param			id path string true "Student ID"
// @Produce		json
// @Success		200 {object} models.Response "Status checked successfully"
// @Failure		400 {object} models.Response "Bad Request"
// @Failure		404 {object} models.Response "Not Found"
// @Failure		500 {object} models.Response "Internal Server Error"
func (h Handler) StudentStatus(c *gin.Context) {
	id := c.Param("id")
	if err := uuid.Validate(id); err != nil {
		handleResponse(c, fmt.Sprintf("Error while validating ID %s", id), http.StatusBadRequest, err.Error())
		fmt.Println(id)
		return
	}

	boolean, err := h.Store.StudentStorage().StatusSt(id)
	if err != nil {
		handleResponse(c, "Error while checking student status", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Status checked successfully", http.StatusOK, boolean)
}
