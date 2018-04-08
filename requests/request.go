//Package requests : core handling of all the http request response calls
package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/errors"
)

//Auth : Struct for handling the client authentication keys
type Auth struct {
	Key    string
	Secret string
}

//TIMEOUT : Default SDK Timeout in seconds
const TIMEOUT = 10

//Request : Basic Struct for holding the entire request/response information for the SDK
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
		parameters.Add(k, fmt.Sprintf("%v", v))
	}

	URL.RawQuery = parameters.Encode()

	return URL.String()
}

//AddHeaders : method to add additonal headers to the request
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

//SetTimeout : method to explicity set client timeout
func (request *Request) SetTimeout(timeout int16) {
	timeoutSeconds := int64(timeout) * int64(time.Second)
	request.HTTPClient = &http.Client{Timeout: time.Duration(timeoutSeconds)}
}

func processResponse(response *http.Response) (map[string]interface{}, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	resp := make(map[string]interface{})
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (request *Request) doRequestResponse(req *http.Request) (map[string]interface{}, error) {

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

	switch errorData.InternalErrorCode {
	case constants.SERVER_ERROR:
		return nil, &errors.ServerError{Message: errorData.Description}
	case constants.GATEWAY_ERROR:
		return nil, &errors.GatewayError{Message: errorData.Description}
	case constants.BAD_REQUEST_ERROR:
	default:
		return nil, &errors.BadRequestError{Message: errorData.Description}
	}

	return processResponse(response)
}

//Get : method to handle HTTP GET calls. Takes the request path, the query params and additional http options as arguments.
func (request *Request) Get(path string, queryParams map[string]interface{}, options map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	url = buildURLWithParams(url, queryParams)

	req, _ := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Post : method to handle HTTP POST calls. Takes the request path, the post payload and additional http options
// to be used as part of the request
func (request *Request) Post(path string, payload map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Patch : method to handle HTTP PATCH calls. Takes request path, patch payload and additional http options to be used as part of the request
func (request *Request) Patch(path string, payload map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Put : method to handle HTTP PUT calls. Takes request path, payload and additional http options to be used as part of the request
func (request *Request) Put(path string, payload map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}

//Delete : method to handle HTTP DELETE calls. Takes request path, query params and additional http options to be used as part of the request
func (request *Request) Delete(path string, queryParams map[string]interface{}, options map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	url = buildURLWithParams(url, queryParams)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, options)

	return request.doRequestResponse(req)
}
