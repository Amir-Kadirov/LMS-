package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router		/timetable [post]
// @Summary		Creates a timetable
// @Description	This api creates a timetable and returns its id
// @Tags		TimeTable
// @Accept		json
// @Produce		json
// @Param		timetable body models.TimeTable true "timetable"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) CreateTimeTable(c *gin.Context) {
	timetable := models.TimeTable{}

	if err := c.ShouldBindJSON(&timetable); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.TimeTable().CreateTimeTable(c.Request.Context(), timetable)
	if err != nil {
		handleResponse(c, h.Log, "error while creating timetable", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Created successfully", http.StatusOK, id)
}

// @Security ApiKeyAuth
// @Router		/timetable/delettimetable/{id} [delete]
// @Summary		Delete a timetable
// @Description	This API delete a timetable
// @Tags		TimeTable
// @Produce		json
// @Param		id path string true "timetable ID"
// @Success		200 {object} models.Response
// @Failure		400 {object} models.Response
// @Failure		404 {object} models.Response
// @Failure		500 {object} models.Response
func (h Handler) DeleteTimeTable(c *gin.Context) {

	Id := c.Param("id")

	err := h.Service.TimeTable().DeleteTimeTable(c.Request.Context(), Id)
	if err != nil {
		handleResponse(c, h.Log, "error while deleting timetable", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Time Table deleted successfully", http.StatusOK, err)
}

// @Security ApiKeyAuth
// @Router		/timetable/studentsattandence [post]
// @Summary		get students attandence
// @Description	This api get students attandence 
// @Tags		TimeTable
// @Accept		json
// @Produce		json
// @Param		timetable body models.GetAllStudentsAttandenceReportRequest true "Attanddence Students"
// @Success		200  {object}  models.Response
// @Failure		400  {object}  models.Response
// @Failure		404  {object}  models.Response
// @Failure		500  {object}  models.Response
func (h Handler) GetAllStudentsAttandenceReport(c *gin.Context) {
	req := models.GetAllStudentsAttandenceReportRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		handleResponse(c, h.Log, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.TimeTable().GetAllStudentsAttandenceReport(c.Request.Context(), req)
	if err != nil {
		handleResponse(c, h.Log, "error while get attandence students", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Getted successfully", http.StatusOK, id)
}
/*
git merch
git workflow
	*/