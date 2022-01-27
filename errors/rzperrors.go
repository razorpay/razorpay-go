package errors

//RZPError ...
type RZPError struct {
	InternalErrorCode string `json:"internal_error_code"`
	Field             string `json:"field"`
	Description       string `json:"description"`
	Code              string `json:"code"`
}

//RZPErrorJSON ...
type RZPErrorJSON struct {
	ErrorData RZPError `json:"error"`
}

//BadRequestError ...
type BadRequestError struct {
	Message string
}

//ServerError ...
type ServerError struct {
	Message string
}

//GatewayError ...
type GatewayError struct {
	Message string
}

//SignatureVerificationError ...
type SignatureVerificationError struct {
	Message string
}

func handleError(msg string) string {
	errorMessage := ""
	if msg != "" {
		errorMessage += msg
	}

	return errorMessage
}

//Error ...
func (s *BadRequestError) Error() string {
	return handleError(s.Message)
}

func (s *GatewayError) Error() string {
	return handleError(s.Message)

}

func (s *ServerError) Error() string {
	return handleError(s.Message)
}

func (s *SignatureVerificationError) Error() string {
	return handleError(s.Message)
}
