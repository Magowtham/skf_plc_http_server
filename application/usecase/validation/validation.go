package validation

import (
	"fmt"
	"regexp"
	"unicode"
)

func ValidatePassword(password string) error {

	var (
		hasUpperCase bool
		hasLowerCase bool
		hasDigit     bool
		hasSpecial   bool
	)

	if len(password) < 8 {
		return fmt.Errorf("password length should be greater then 8")
	}

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpperCase = true
		} else if unicode.IsLower(char) {
			hasLowerCase = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else {
			hasSpecial = true
		}
	}

	if !hasUpperCase {
		return fmt.Errorf("password should contain atleast 1 upper case letter")
	} else if !hasLowerCase {
		return fmt.Errorf("password should contain atleast 1 lower case letter")
	} else if !hasDigit {
		return fmt.Errorf("password should contain atleast 1 numeric digit")
	} else if !hasSpecial {
		return fmt.Errorf("password should contain atleast 1 special character")
	} else {
		return nil
	}
}

func ValidateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	re := regexp.MustCompile(emailRegex)

	if !re.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}

	return nil
}
