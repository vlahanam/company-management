package common

type successResponse struct {
	Message    string       `json:"message"`
	Data       *interface{} `json:"data,omitempty"`
	Pagination *Pagination  `json:"pagination,omitempty"`
	Query      *interface{} `json:"query,omitempty"`
}

func NewSuccessResponse(
	message string,
	pagination Pagination,
	data, query interface{},
) *successResponse {
	return &successResponse{
		Message:    message,
		Data:       ptr(data),
		Pagination: &pagination,
		Query:      ptr(query),
	}
}

func (sr *successResponse) WrapMessage(message string) *successResponse {
	sr.Message = message
	return sr
}

func (sr *successResponse) WrapPagination(pagination Pagination) *successResponse {
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
	Message: "Login successful",
}

var RegisterSuccessful = &successResponse{
	Message: "Registration successful",
}

var CreatedSuccessfully = &successResponse{
	Message: "Created successfully",
}

var UpdatedSuccessfully = &successResponse{
	Message: "Updated successfully",
}

var DeletedSuccessfully = &successResponse{
	Message: "Deleted successfully",
}

var RetrievedSuccessfully = &successResponse{
	Message: "Retrieved successfully",
}
