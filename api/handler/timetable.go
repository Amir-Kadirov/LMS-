package handler

import (
	_ "backend_course/lms/api/docs"
	"backend_course/lms/api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
	timetable:=models.TimeTable{}

	if err := c.ShouldBindJSON(&timetable); err != nil {
		handleResponse(c, "error while reading request body", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.Service.TimeTable().CreateTimeTable(c.Request.Context(),timetable)
	if err != nil {
		handleResponse(c, "error while creating timetable", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, "Created successfully", http.StatusOK, id)
}

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

	Id:= c.Param("id")

	err := h.Service.TimeTable().DeleteTimeTable(c.Request.Context(),Id)
	if err != nil {
		handleResponse(c, "error while deleting timetable", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, "Time Table deleted successfully", http.StatusOK, err)
}