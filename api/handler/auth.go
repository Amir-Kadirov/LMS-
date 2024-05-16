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
	loginReq := models.LoginRequest{}

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

	loginResp, err := h.Service.Auth().TeacherLogin(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, loginResp)

}

// TeacherRegister godoc
// @Router       /teacher/register [POST]
// @Summary      Teacher register
// @Description  Teacher register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        register body models.RegisterRequest true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherRegister(c *gin.Context) {
	loginReq := models.RegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	if !check.ValidateGmail(loginReq.Mail) {
		handleResponse(c, h.Log, "error while validating email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	err := h.Service.Auth().TeacherRegister(c.Request.Context(), loginReq)
	if err != nil {
		handleResponse(c, h.Log, "Bad request", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, "Your request succeed")

}

// TeacherRegister godoc
// @Router       /teacher/register-confirm [POST]
// @Summary      Teacher register-confirm
// @Description  Teacher register-confirm
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        teacher body models.TeacherRegisterConfirm true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherRegisterConfirm(c *gin.Context) {
	teacher := models.TeacherRegisterConfirm{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	if !check.ValidateGmail(teacher.Teacher.Mail) {
		handleResponse(c, h.Log, "error while validating email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	id, err := h.Service.Auth().TeacherRegisterComfirm(c.Request.Context(), teacher)
	if err != nil {
		handleResponse(c, h.Log, "Bad request", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, id)

}

// TeacherLogin godoc
// @Router       /teacher/loginbymail [POST]
// @Summary      Teacher login
// @Description  Teacher login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.RegisterRequest true "login"
// @Success      201  {object}  models.LoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherLoginByMail(c *gin.Context) {
	loginReq := models.RegisterRequest{}

	if err := c.ShouldBindJSON(&loginReq); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}
	fmt.Println("loginReq: ", loginReq)

	if !check.ValidateGmail(loginReq.Mail) {
		handleResponse(c, h.Log, "error while validation email", http.StatusBadRequest, errors.New("wrong email"))
		return
	}

	err := h.Service.Auth().TeacherLoginByMail(c.Request.Context(),loginReq)
	if err != nil {
		handleResponse(c, h.Log, "unauthorized", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, nil)

}


// TeacherRegister godoc
// @Router       /teacher/login-confirm [POST]
// @Summary      Teacher login-confirm
// @Description  Teacher login-confirm
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        teacher body models.LoginRequestEmail true "register"
// @Success      201  {object}  models.Response
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h *Handler) TeacherLoginConfirm(c *gin.Context) {
	teacher := models.LoginRequestEmail{}

	if err := c.ShouldBindJSON(&teacher); err != nil {
		handleResponse(c, h.Log, "error while binding body", http.StatusBadRequest, err)
		return
	}

	resp, err := h.Service.Auth().TeacherLoginByMailConfirm(c.Request.Context(),teacher)
	if err != nil {
		handleResponse(c, h.Log, "Bad request", http.StatusInternalServerError, err.Error())
		return
	}

	handleResponse(c, h.Log, "Succes", http.StatusOK, resp)

}
