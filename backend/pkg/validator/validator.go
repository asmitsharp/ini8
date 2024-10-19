package validator

import (
	"fmt"
	"ini8/internal/model"
)

func ValidateRegistration(reg *model.Registration) error {
	if reg.Name == "" {
		return fmt.Errorf("name is required")
	}

	if reg.Email == "" {
		return fmt.Errorf("email is required")
	}

	// if reg.DOB.After(time.Now()) {
	// 	return fmt.Errorf("date of birth cannot be in the future")
	// }

	return nil
}
