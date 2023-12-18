package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/errors"
)

//Auth holds the values required to authenticate the requests made to Razorpay APIs
type Auth struct {
	Key    string
	Secret string
}

type FileUploadParams struct {
	File   *os.File
	Fields map[string]string
}

//TIMEOUT ... client timeout
const TIMEOUT = 10

// Request encapsulates all the information required to make an HTTP request
// to Razorpay's APIs.
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

// AddHeaders adds header to the Request object
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

func (request *Request) addRequestHeaders(req *http.Request, headers map[string]string, contentType ...string) {
	//Set the Defaults First in case unavailable
	req.Header.Set("User-Agent", fmt.Sprintf("%s/%s", request.SDKName, request.Version))

	if len(contentType) == 0 {
		req.Header.Set("Content-Type", "application/json")
	} else {
		req.Header.Set("Content-Type", contentType[0])
	}

	// Set the already added headers
	request.addRequestHeadersInternal(req, request.Headers)
	request.addRequestHeadersInternal(req, headers)
}

//SetTimeout ...
func (request *Request) SetTimeout(timeout int16) {
	timeoutSeconds := int64(timeout) * int64(time.Second)
	request.HTTPClient = &http.Client{Timeout: time.Duration(timeoutSeconds)}
}

func processResponse(response *http.Response) (map[string]interface{}, error) {
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if len(body) == 0 || len(body) == 2 {
		resp := make(map[string]interface{})
		return resp, nil
	}

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
	json.NewDecoder(response.Body).Decode(&jsonResponse)
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

//Get ...
func (request *Request) Get(path string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	url = buildURLWithParams(url, queryParams)

	req, _ := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, extraHeaders)

	return request.doRequestResponse(req)
}

//Post ...
func (request *Request) Post(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, extraHeaders)

	return request.doRequestResponse(req)
}

//Patch ...
func (request *Request) Patch(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, extraHeaders)

	return request.doRequestResponse(req)
}

//Put ...
func (request *Request) Put(path string, payload map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	jsonStr, _ := json.Marshal(payload)

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, _ := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, extraHeaders)

	return request.doRequestResponse(req)
}

//Delete ...
func (request *Request) Delete(path string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	url = buildURLWithParams(url, queryParams)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, extraHeaders)

	return request.doRequestResponse(req)
}

func (request *Request) File(path string, params FileUploadParams, extraHeaders map[string]string) (map[string]interface{}, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add the file field to the multipart form data
	filePart, err := writer.CreateFormFile("file", filepath.Base(params.File.Name()))
	if err != nil {
		log.Fatal(err)
	}
	// Copy the file contents to the form file part
	_, err = io.Copy(filePart, params.File)
	if err != nil {
		log.Fatal(err)
	}
	// Add additional form fields
	for fieldName, fieldValue := range params.Fields {
		err = writer.WriteField(fieldName, fieldValue)
		if err != nil {
			log.Fatal(err)
		}
	}
	// Close the writer to finalize the form data
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}
	// Create the HTTP request
	url := fmt.Sprintf("%s%s", request.BaseURL, path)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Fatal(err)
	}
	contentType := writer.FormDataContentType()

	req.SetBasicAuth(request.Auth.Key, request.Auth.Secret)

	request.addRequestHeaders(req, extraHeaders, contentType)

	return request.doRequestResponse(req)
}
