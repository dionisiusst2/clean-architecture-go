package validator

import (
	"fmt"
	"regexp"
)

func ValidateNotEmpty(fieldName string, value string) string {
	if len(value) == 0 {
		return fmt.Sprintf("%v cannot be empty", fieldName)
	}
	return ""
}

func ValidateLengthBetween(fieldName string, value string, min int, max int) string {
	if len(value) < min || len(value) > max {
		return fmt.Sprintf("%v must be %v-%v characters", fieldName, min, max)
	}
	return ""
}

func ValidateEmailFormat(email string) string {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if ok := emailRegex.MatchString(email); !ok {
		return "email format is incorrect"
	}
	return ""
}
