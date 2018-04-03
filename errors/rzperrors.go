package errors

// Types of Razorpay Errors
const (
	BAD_REQUEST_ERROR = "BAD_REQUEST_ERROR"
	SERVER_ERROR      = "SERVER_ERROR"
	GATEWAY_ERROR     = "GATEWAY_ERROR"
)

//RZPError ...
type RZPError struct {
	InternalErrorCode string `json:"internal_error_code"`
	Field             string `json:"field"`
	Description       string `json:"description"`
	Code              string `json:"code"`
}

//GetInternalErrorCode ...
func (e *RZPError) GetInternalErrorCode() string {
	return e.InternalErrorCode
}

//GetDescription ...
func (e *RZPError) GetDescription() string {
	return e.Description
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

//Error ...
func (s *BadRequestError) Error() string {
	errorMessage := ""

	if s == nil {
		return "<nil>"
	}

	if s.Message != "" {
		errorMessage += s.Message
	}

	errorMessage += s.Err.Error()

	return errorMessage
}

func handleError(msg string, desc string) string {
	errorMessage := ""
	if msg != "" {
		errorMessage += msg
	}
	errorMessage += desc
	return errorMessage
}

func (s *GatewayError) Error() string {

	if s == nil {
		return "<nil>"
	}

	return handleError(s.Message, s.Err.Error())

}

func (s *ServerError) Error() string {
	if s == nil {
		return "<nil>"
	}

	return handleError(s.Message, s.Err.Error())
}

func (s *SignatureVerificationError) Error() string {
	if s == nil {
		return "<nil>"
	}

	return handleError(s.Message, s.Err.Error())

}
