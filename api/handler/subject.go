package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.Subject().CreateSubject(subject)
	if err != nil {
		handleResponse(c, "error while creating subject", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

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

	data, err := h.Service.Subject().GetbyIdSubject(id)
	if err != nil {
		handleResponse(c, fmt.Sprintf("error while get by id subject %s", id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Select by id successful", http.StatusOK, data)
}

// @Router		/subject/updatesubject/{id} [put]
// @Summary		Update a subject
// @Description	This API updates a Teacher
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

	subject:=models.Subjects{}

	if err := c.ShouldBindJSON(&subject); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	subject.Id = id

	err := h.Service.Subject().UpdateSubject(subject)
	if err != nil {
		handleResponse(c, fmt.Sprintf("error while updating subject %s", subject.Id), http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Updated successfully", http.StatusOK, id)
}

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

	Id:= c.Param("id")

	err := h.Service.Subject().DeleteSubject(Id)
	if err != nil {
		handleResponse(c, "error while deleting  Subject", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Subject deleted successfully", http.StatusOK, err)
}