package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router		/subject [post]
// @Summary		Creates a subject
// @Description	This api creates a subject and returns its id
// @Tags		Subject
// @Accept		json
// @Produce		json
// @Param		subject body models.Subjects true "subject"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateSubject(c *gin.Context) {
	subject := models.Subjects{}

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Subject().CreateSubject(c.Request.Context(), subject)
	if err != nil {
		handleResponse(c, h.Log, "error while creating subject", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/subject/{id} [get]
// @Summary		Get by id a subject
// @Description	This API get by id a subject
// @Tags		Subject
// @Produce		json
// @Param		id path string true "Subject id"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) GetbyIdSubject(c *gin.Context) {
	id := c.Param("id")

	data, err := h.Service.Subject().GetbyIdSubject(c.Request.Context(), id)
	if err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("error while get by id subject %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Select by id successful", http.StatusOK, data)
}

// @Security ApiKeyAuth
// @Router		/subject/updatesubject/{id} [put]
// @Summary		Update a subject
// @Description	This API updates a Subject
// @Tags		Subject
// @Accept		json
// @Produce		json
// @Param		id path string true "Subject ID"
// @Param		subject body models.Subjects true "subject object to update"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) UpdateSubject(c *gin.Context) {

	id := c.Param("id")

	subject := models.Subjects{}

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	subject.Id = id

	err := h.Service.Subject().UpdateSubject(c.Request.Context(), subject)
	if err != nil {
		handleResponse(c, h.Log, fmt.Sprintf("error while updating subject %s", subject.Id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Updated successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/subject/deletsubject/{id} [delete]
// @Summary		Delete a Subject
// @Description	This API delete a Subject
// @Tags		Subject
// @Produce		json
// @Param		id path string true "Subject ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) DeleteSubject(c *gin.Context) {

	Id := c.Param("id")

	err := h.Service.Subject().DeleteSubject(c.Request.Context(), Id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting  Subject", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Subject deleted successfully", http.StatusOK, err)
}

// @Security ApiKeyAuth
// @Router		/subject [get]
// @Summary		Get a Subject
// @Description	This API returns all Subject
// @Tags		Subject
// @Produce		json
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllSubject(c *gin.Context) {
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

	resp, err := h.Service.Subject().GetAllSubject(c.Request.Context(), models.GetAllStudentsRequest{
		Search: search,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		handleResponse(c, h.Log, "error while getting all Subject", http.StatusInternalServerError, err.Error())
		return
	}
	handleResponse(c, h.Log, "Select all successful", http.StatusOK, resp)
}
