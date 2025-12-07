package common

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/vlahanam/company-management/utils"
)

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

func (e *rootError) WrapError(err error) *rootError {
	e.Err = err
	return e
}

func (e *rootError) WrapErrorSafe(err error) *rootError {
	e.Err = err

	if utils.IsDevelopment() {
		debug := err.Error()
		e.Debug = &debug
	}

	return e
}

func (e *rootError) SetDetail(key, value string) *rootError {
	detail := map[string]string{key: value}
	e.Detail = &detail
	return e
}

func (e *rootError) Error() string {
	return e.Message
}

func (e *rootError) Unwrap() error {
	return e.Err
}

func (e *rootError) Clone() *rootError {
	clone := &rootError{
		Key:     e.Key,
		Message: e.Message,
		Err:     e.Err,
	}

	if e.Debug != nil {
		debug := *e.Debug
		clone.Debug = &debug
	}

	if e.Detail != nil {
		detail := make(map[string]string)
		for k, v := range *e.Detail {
			detail[k] = v
		}
		clone.Detail = &detail
	}

	return clone
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
var ErrorQueryParser = &rootError{
	Key:     "QUERY_PARSER_ERROR",
	Message: "invalid request query format",
}
var ErrorCreateFailed = &rootError{
	Key:     "CREATE_FAILED",
	Message: "failed to create resource",
}
