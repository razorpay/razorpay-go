package resources

import (
	"errors"
	"fmt"
	"log"

	"github.com/razorpay/razorpay-go/requests"
)

type Entity struct {
	Request *requests.Request
}

func (e *Entity) Do(endpoint string, method string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {

	url := fmt.Sprintf("/%s", endpoint)

	if requests.IsEntityExist(endpoint) {
		log.Println("Warning: The entity already has a specific function. Consider using it instead.")
	}

	switch method {
	case "GET":
		return e.Request.Get(url, data, extraHeaders)
	case "POST":
		return e.Request.Post(url, data, extraHeaders)
	case "PUT":
		return e.Request.Put(url, data, extraHeaders)
	case "PATCH":
		return e.Request.Patch(url, data, extraHeaders)
	case "DELETE":
		return e.Request.Delete(url, data, extraHeaders)
	default:

	}

	return nil, errors.New("Unsupported method or error occurred")
}
