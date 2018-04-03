package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"gitlab.com/venkatvghub/razorpay-go/razorpay/constants"
	"gitlab.com/venkatvghub/razorpay-go/razorpay/errors"
)

//Auth ...
type Auth struct {
	Key    string
	Secret string
}

//TIMEOUT ... client timeout
const TIMEOUT = 10

//Request ...
type Request struct {
	Auth       Auth
	HTTPClient *http.Client
	Headers    map[string]string
	Version    string
	SDKName    string
	AppDetails map[string]string
	BaseURL    string
}

func buildURLWithParams(requestURL string, data map[string]interface{}) string {
	var URL *url.URL

	URL, err := url.Parse(requestURL)

	if err != nil {
		panic(err)
	}

	parameters := url.Values{}

	for k, v := range data {
		parameters.Add(k, v.(string))
	}

	URL.RawQuery = parameters.Encode()

	return URL.String()
}

//AddHeaders ...
func (request *Request) AddHeaders(headers map[string]string) {
	//This just appends the headers to the request struct. In the actual
	// Get/Post/Put/Delete calls, we call `addRequestHeaders` method
	for key, value := range headers {
		request.Headers[key] = value
	}
}

func (request *Request) addRequestHeadersInternal(req *http.Request, headers map[string]string) {
	for key, value := range headers {
		if key == "Content-Type" || key == "User-Agent" {
			continue
		}
		req.Header.Set(key, value)
	}
}

func (request *Request) addRequestHeaders(req *http.Request, headers map[string]string) {
	//Set the Defaults First in case unavailable
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", request.SDKName, request.Version))
	req.Header.Set("Content-Type", "application/json")

	// Set the already added headers
	request.addRequestHeadersInternal(req, request.Headers)
	request.addRequestHeadersInternal(req, headers)
}

//SetTimeout ...
func (request *Request) SetTimeout(timeout int16) {
	timeoutSeconds := int64(timeout) * int64(time.Second)
	request.HTTPClient = &http.Client{Timeout: time.Duration(timeoutSeconds)}
}

func processResponse(response *http.Response) ([]byte, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (request *Request) doRequestResponse(req *http.Request) ([]byte, error) {

	client := request.HTTPClient

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	if response.StatusCode >= constants.HTTP_STATUS_OK &&
		response.StatusCode < constants.HTTP_STATUS_REDIRECT {
		return processResponse(response)
	}

	// Raise an error depending on the type of error in response
	var jsonResponse errors.RZPErrorJSON
	json.NewDecoder(response.Body).Decode(jsonResponse)
	errorData := jsonResponse.ErrorData

	switch errorData.GetInternalErrorCode() {
	case errors.SERVER_ERROR:
		return nil, &errors.ServerError{Message: errorData.GetDescription()}
	case errors.GATEWAY_ERROR:
		return nil, &errors.GatewayError{Message: errorData.GetDescription()}
	case errors.BAD_REQUEST_ERROR:
	default:
		return nil, &errors.BadRequestError{Message: errorData.GetDescription()}
	}

	return processResponse(response)
}

//Get ...
func (request *Request) Get(path string, queryParams map[string]interface{}, options map[string]string) ([]byte, error) {
	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	url = buildURLWithParams(url, queryParams)

	req, _ := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Post ...
func (request *Request) Post(path string, payload map[string]interface{}, options map[string]string) ([]byte, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	req.Header.Set("Content-Type", "application/json")

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Put ...
func (request *Request) Put(path string, payload map[string]interface{}, options map[string]string) ([]byte, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	req.Header.Set("Content-Type", "application/json")

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Delete ...
func (request *Request) Delete(path string, options map[string]string) ([]byte, error) {

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}
