package common

import (
	"fmt"
	"strings"
)

type successResponse struct {
	Key        string       `json:"key"`
	Message    string       `json:"message"`
	Data       *interface{} `json:"data,omitempty"`
	Pagination *Paging      `json:"pagination,omitempty"`
	Query      *interface{} `json:"query,omitempty"`
}

func NewSuccessFullResponse(key, message string, pagination Paging, data, query interface{}) *successResponse {
	return &successResponse{
		Key:        key,
		Message:    message,
		Data:       ptr(data),
		Pagination: &pagination,
		Query:      ptr(query),
	}
}

func NewSuccessResponse(key, message string) *successResponse {
	return &successResponse{Key: key, Message: message}
}

func CreateSuccessResponse(resourceName string) *successResponse {
	return NewSuccessResponse(
		fmt.Sprintf("CREATED_%s_SUCCESS", strings.ToUpper(resourceName)),
		fmt.Sprintf("%s created successfully", resourceName),
	)
}

func UpdateSuccessResponse(resourceName string) *successResponse {
	return NewSuccessResponse(
		fmt.Sprintf("UPDATED_%s_SUCCESS", strings.ToUpper(resourceName)),
		fmt.Sprintf("%s updated successfully", resourceName),
	)
}

func DeleteSuccessResponse(resourceName string) *successResponse {
	return NewSuccessResponse(
		fmt.Sprintf("DELETED_%s_SUCCESS", strings.ToUpper(resourceName)),
		fmt.Sprintf("%s deleted successfully", resourceName),
	)
}

func GetSuccessResponse(resourceName string) *successResponse {
	return NewSuccessResponse(
		fmt.Sprintf("GET_%s_SUCCESS", strings.ToUpper(resourceName)),
		fmt.Sprintf("Get %s successfully", resourceName),
	)
}

func GetListSuccessResponse(resourceName string) *successResponse {
	return NewSuccessResponse(
		fmt.Sprintf("GET_LIST_%s_SUCCESS", strings.ToUpper(resourceName)),
		fmt.Sprintf("Get list %s successfully", resourceName),
	)
}

func (sr *successResponse) WrapMessage(message string) *successResponse {
	sr.Message = message
	return sr
}

func (sr *successResponse) WrapPagination(pagination Paging) *successResponse {
	sr.Pagination = &pagination
	return sr
}

func (sr *successResponse) WrapData(data interface{}) *successResponse {
	sr.Data = &data
	return sr
}

func (sr *successResponse) WrapQuery(query interface{}) *successResponse {
	sr.Query = &query
	return sr
}

func ptr[T any](v T) *T {
	return &v
}

var LoginSuccessful = &successResponse{
	Key: "LOGIN_SUCCESSFUL",
	Message: "Login successful",
}
var RegisterSuccessful = &successResponse{
	Message: "Registration successful",
}
var RetrievedSuccessfully = &successResponse{
	Message: "Retrieved successfully",
}
var AccessTokenCreatedSuccessfully = &successResponse{
	Message: "Access token created successfully",
}
