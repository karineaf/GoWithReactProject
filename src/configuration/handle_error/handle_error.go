package handle_error

import "net/http"

type HandleError struct {
	Message string 		`json:"message"`
	Err 	string 		`json:"error"` 
	Code 	int			`json:"code"` 
	Causes 	[]Causes	`json:"causes"` 
}

type Causes struct {
	Field 	string	`json:"field"` 
	Message string	`json:"message"` 
}

func(r *HandleError) Error() string {
	return r.Message
}

func HandlerError(message, err string, code int, causes []Causes) *HandleError{
	return &HandleError{
		Message: message,
		Err: err, 
		Code: code,
		Causes: causes,
	}
}

func BadRequestError(message string) *HandleError{
	return &HandleError{
		Message: message,
		Err: "Bad Request", 
		Code: http.StatusBadRequest, 	}
}

func BadRequestValidationError(message string, causes []Causes) *HandleError{
	return &HandleError{
		Message: message,
		Err: "bad_request", 
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func InternalServerError(message string) *HandleError{
	return &HandleError{
		Message: message,
		Err: "internal_server_error", 
		Code: http.StatusInternalServerError,
	}
}

func NotFoundError(message string) *HandleError{
	return &HandleError{
		Message: message,
		Err: "not_found", 
		Code: http.StatusNotFound,
	}
}

func ForbiddenError(message string) *HandleError{
	return &HandleError{
		Message: message,
		Err: "forbidden", 
		Code: http.StatusForbidden,
	}
}