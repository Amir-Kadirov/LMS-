package handler

import (
	"backend_course/lms/api/models"
	"backend_course/lms/pkg/check"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TeacherLogin godoc
// @Router       /teacher/login [POST]
// @Summary      Teacher login
// @Description  Teacher login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.LoginRequest true "login"
// @Success      201  {object}  models.LoginResponse
// @Failure      400  {object}  models.Response

// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherLogin(c *gin.Context) {
	loginReq:=models.LoginRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	if !check.ValidateGmail(loginReq.Login) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	if !check.ValidatePassword(loginReq.Password) {
		handleResponse(c, h.Log, "error while validating teacher password", http.StatusBadRequest, errors.New("unsecure password"))
		return
	}

	loginResp,err:=h.Service.Auth().TeacherLogin(c.Request.Context(),loginReq)
	if err!=nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}