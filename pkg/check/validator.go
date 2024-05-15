package check

import (
	"errors"
	"regexp"
	"time"
)

func ValidateAge(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}

	return nil
}

func ValidatePhone(phone string) bool {
	return regexp.MustCompile(`^\+998[0-9]{9}$`).MatchString(phone)
}

func ValidateGmail(gmail string) bool {
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@gmail.com$`).MatchString(gmail)
}

func ValidatePassword(password string) bool {
	return regexp.MustCompile(`^[A-Z0-9!@#$+]{9}`).MatchString(password)
}