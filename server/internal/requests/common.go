package requests

import (
	"net/mail"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func isValidEmail() validation.Rule {
	return validation.By(func(value interface{}) error {
		email, _ := value.(string)

		_, err := mail.ParseAddress(email)
		if err != nil {
			return validation.NewError("validation_invalid_email", "Invalid email")
		}

		return nil
	})
}

func FormatValidationError(err error) map[string]any {
	if errs, ok := err.(validation.Errors); ok {
		return map[string]interface{}{"detail": errs}
	}
	return map[string]interface{}{"detail": map[string]string{"error": err.Error()}}
}
