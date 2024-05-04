package check

import (
	"errors"
	"time"
)

func ValidateAge(year int) error {
	if year <= 0 || year > time.Now().Year()+1 {
		return errors.New("year is not valid")
	}

	return nil
}

func ValidateNumber(num string) error {
	if len(num) != 13 || num[:4] != "+998" {
		return errors.New("phone is not valid")
	}
	return nil
}