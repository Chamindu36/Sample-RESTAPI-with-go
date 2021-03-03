package errors

//ServiceError is used to return error messages in business logic
type ServiceError struct{
	Message string `json:"message"`
}
