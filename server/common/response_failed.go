package common

import validation "github.com/go-ozzo/ozzo-validation/v4"

type rootError struct {
	Key     string             `json:"key"`
	Message string             `json:"message"`
	Debug   *string            `json:"debug,omitempty"`
	Detail  *map[string]string `json:"detail,omitempty"`
	Err     error              `json:"-"`
}

func NewRootError(
	key, message, debug string,
	detail map[string]string,
	err error,
) *rootError {
	return &rootError{
		Key:     key,
		Message: message,
		Debug:   &debug,
		Detail:  &detail,
		Err:     err,
	}
}

func (e *rootError) WrapKey(key string) *rootError {
	e.Key = key
	return e
}

func (e *rootError) WrapMessage(message string) *rootError {
	e.Message = message
	return e
}

func (e *rootError) WrapDebug(debug string) *rootError {
	e.Debug = &debug
	return e
}

// Validation error
func (e *rootError) WrapDetail(err error) *rootError {
	detail := make(map[string]string)
	if errs, ok := err.(validation.Errors); ok {
		for field, ferr := range errs {
			if ferr != nil {
				detail[field] = ferr.Error()
			}
		}
	}

	e.Detail = &detail
	return e
}

func (e *rootError) SetDetail(key, value string) *rootError {
	detail := map[string]string{key: value}
	e.Detail = &detail
	return e
}

func (e *rootError) WrapError(err error) *rootError {
	e.Err = err
	return e
}

func (e *rootError) Unwrap() error {
	return e.Err
}

var ErrorValidation = &rootError{
	Key:     "VALIDATION_ERROR",
	Message: "validation error",
}
var ErrorNotFound = &rootError{
	Key:     "NOT_FOUND_ERROR",
	Message: "resource not found",
}
var ErrorUnauthorized = &rootError{
	Key:     "UNAUTHORIZED_ERROR",
	Message: "unauthorized access",
}
var ErrorInternal = &rootError{
	Key:     "INTERNAL_ERROR",
	Message: "internal server error",
}
var ErrorBodyParser = &rootError{
	Key:     "BODY_PARSER_ERROR",
	Message: "invalid request body format",
}
var ErrorCreateFailed = &rootError{
	Key:     "CREATE_FAILED",
	Message: "failed to create resource",
}
