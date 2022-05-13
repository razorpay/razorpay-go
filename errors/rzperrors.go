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
	Err     error
}

//ServerError ...
type ServerError struct {
	Message string
	Err     error
}

//GatewayError ...
type GatewayError struct {
	Message string
	Err     error
}

//SignatureVerificationError ...
type SignatureVerificationError struct {
	Message string
	Err     error
}

func handleError(msg string, err error) string {
	errorMessage := ""
	if msg != "" {
		errorMessage += msg
	}
	
	if err != nil {
		errorMessage += err.Error()
	}

	return errorMessage
}

//Error ...
func (s *BadRequestError) Error() string {
	return handleError(s.Message, s.Err)
}

func (s *GatewayError) Error() string {
	return handleError(s.Message, s.Err)

}

func (s *ServerError) Error() string {
	return handleError(s.Message, s.Err)
}

func (s *SignatureVerificationError) Error() string {
	return handleError(s.Message, s.Err)
}
