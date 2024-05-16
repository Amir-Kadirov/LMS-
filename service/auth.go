package service

import (
	"backend_course/lms/api/models"
	"backend_course/lms/config"
	"backend_course/lms/pkg"
	"backend_course/lms/pkg/hash"
	smtp "backend_course/lms/pkg/helper"
	"backend_course/lms/pkg/jwt"
	"backend_course/lms/pkg/logger"
	"backend_course/lms/storage"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type authService struct {
	storage storage.IStorage
	logger  logger.ILogger
}

func NewAuthService(storage storage.IStorage, logger logger.ILogger) authService {
	return authService{
		storage: storage,
		logger:  logger,
	}
}

func (s authService) TeacherLogin(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {
	resp := models.LoginResponse{}

	teacher, err := s.storage.TeacherStorage().GetTeacherbyLogin(ctx, req.Login)
	if err != nil {
		return resp, err
	}

	if err = hash.CompareHashAndPassword(teacher.Password, req.Password); err != nil {
		return resp, errors.New("password doesn't matched")
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = teacher.Id
	m["user_role"] = config.TEACHER_TYPE
	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}

func (s authService) TeacherRegister(ctx context.Context, req models.RegisterRequest) error {
	_, err := s.storage.TeacherStorage().GetTeacherbyLogin(ctx, req.Mail)
	if err == pgx.ErrNoRows {
		otp := pkg.GenerateOTP()
		msg := fmt.Sprintf("Your OTP: %v. DON'T give anyone", otp)
		err := s.storage.Redis().SetX(ctx, req.Mail, otp, time.Minute*2)
		if err != nil {
			return err
		}

		err = smtp.SendMail(req.Mail, msg)
		if err != nil {
			return err
		}

	} else if err != nil {
		return err
	} else {
		return errors.New("email already exists")
	}

	return nil
}

func (s authService) TeacherRegisterComfirm(ctx context.Context, teacher models.TeacherRegisterConfirm) (string, error) {

	validotp := s.storage.Redis().Get(ctx, teacher.Teacher.Mail)
	if teacher.Otp != validotp {
		return "", errors.New("wrong otp")
	}

	id, err := s.storage.TeacherStorage().CreateTeacher(ctx, teacher.Teacher)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (s authService) TeacherLoginByMail(ctx context.Context, req models.RegisterRequest) error {
	_, err := s.storage.TeacherStorage().GetTeacherbyLogin(ctx, req.Mail)
	if err != pgx.ErrNoRows {
		otp := pkg.GenerateOTP()
		msg := fmt.Sprintf("Your OTP: %v. DON'T give anyone", otp)
		err := s.storage.Redis().SetX(ctx, req.Mail, otp, time.Minute*3)
		if err != nil {
			return err
		}
		err = smtp.SendMail(req.Mail, msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s authService) TeacherLoginByMailConfirm(ctx context.Context, req models.LoginRequestEmail) (models.LoginResponse,error) {
	resp:=models.LoginResponse{}
	teacher,err:=s.storage.TeacherStorage().GetTeacherbyLogin(ctx,req.Login)
	if err!=nil {
		return resp,errors.New("teacher are not exists")
	}

	validotp:=s.storage.Redis().Get(ctx,req.Login)
	if validotp!=req.Otp {
		return resp,errors.New("wrong otp")
	}

	m := make(map[interface{}]interface{})
	m["user_id"] = teacher.Id
	m["user_role"] = config.TEACHER_TYPE
	accessToken, refreshToken, err := jwt.GenJWT(m)
	if err != nil {
		return resp, err
	}
	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	return resp, nil
}
