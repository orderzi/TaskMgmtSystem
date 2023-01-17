package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"
)

func ValidateDate(date string) (string, error) {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	return date, err
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func FormatName(name string) (string, error) {
	name = strings.TrimSpace(name)
	reg, err := regexp.Compile("^[a-zA-Z]+$")
	if err != nil {
		return "", fmt.Errorf("error compiling regex", err)
	}
	if reg.MatchString(name) {
		return strings.Title(name), nil
	} else {
		return "", fmt.Errorf("name '%v' should contain only letters", name)
	}
}

func ValidatePriority(priority string) error {
	elements := [4]string{"critical", "high", "medium", "low"}
	for _, element := range elements {
		if element == strings.ToLower(priority) {
			return nil
		}
	}
	return fmt.Errorf("priority value cannot be '%s'", priority)
}
