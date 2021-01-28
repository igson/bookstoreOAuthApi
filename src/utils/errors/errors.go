package errors

import "net/http"

type RestErroAPI struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
}

//NewBadRequestError retorno das mensages de erro conforme o padrão rest
func NewBadRequestError(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusBadRequest,
		Error:      "bad_request",
	}
}

//NewInternalServerError retorno das mensages de erro conforme o padrão rest
func NewInternalServerError(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusInternalServerError,
		Error:      "internal_server_error",
	}
}

//NewNotFoundErro retorno das mensages de erro conforme o padrão rest
func NewNotFoundErro(message string) *RestErroAPI {
	return &RestErroAPI{
		Message:    message,
		StatusCode: http.StatusNotFound,
		Error:      "not_found",
	}
}
