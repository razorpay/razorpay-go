//Package errors : handled all Razorpay SDK Errors
package errors

//RZPError : Base Error struct for unmarshaling the error json from the Razorpay API
type RZPError struct {
	InternalErrorCode string `json:"internal_error_code"`
	Field             string `json:"field"`
	Description       string `json:"description"`
	Code              string `json:"code"`
}

//RZPErrorJSON : Wrapper over the RZPError struct that actually unmarhsalls the error json respnse
type RZPErrorJSON struct {
	ErrorData RZPError `json:"error"`
}

//BadRequestError : Struct to handle errors of the type `BadRequestError`
type BadRequestError struct {
	Message string
	Err     error
}

//ServerError : Struct to handle errors of the type `ServerError`
type ServerError struct {
	Message string
	Err     error
}

//GatewayError : Struct to handle errors of the type `GatewayError`
type GatewayError struct {
	Message string
	Err     error
}

//SignatureVerificationError : Struct to handle errors of the type `SignatureVerificationError`
type SignatureVerificationError struct {
	Message string
	Err     error
}

func handleError(msg string, desc string) string {
	errorMessage := ""
	if msg != "" {
		errorMessage += msg
	}
	errorMessage += desc
	return errorMessage
}

//Error : interface method for handling BadRequestError
func (s *BadRequestError) Error() string {
	return handleError(s.Message, s.Err.Error())
}

//Error : interface method for handling GatewayError
func (s *GatewayError) Error() string {
	return handleError(s.Message, s.Err.Error())

}

//Error : interface method for handling ServerError
func (s *ServerError) Error() string {
	return handleError(s.Message, s.Err.Error())
}

//Error : interface method for handling SignatureVerificationError
func (s *SignatureVerificationError) Error() string {
	return handleError(s.Message, s.Err.Error())
}
