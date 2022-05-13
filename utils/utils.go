package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
        "crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	razorpay "github.com/razorpay/razorpay-go"
	"github.com/stretchr/testify/assert"
)

var (
	//Mux ...
	mux *http.ServeMux
	//Server ...
	server *httptest.Server
	//Client ...
	Client *razorpay.Client
)

const (
	//TestAccessKey ...
	TestAccessKey = "fake_key"
	//TestAccessSecret ...
	TestAccessSecret = "fake_secret"
)

func testSetup() func() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
	Client = razorpay.NewClient(TestAccessKey, TestAccessSecret)
	razorpay.Request.BaseURL = server.URL

	return func() {
		server.Close()
	}
}

func getFixture(path string) string {
	filepath := path + ".json"
	b, err := ioutil.ReadFile("../testdata/" + filepath)
	if err != nil {
		panic(err)
	}
	// Remove newlines and extra spaces
	return strings.Replace(strings.Replace(string(b), "\n", "", -1), " ", "", -1)
}

func jsonCompare(b1, b2 []byte) (bool, error) {
	var o1, o2 interface{}
	err := json.Unmarshal(b1, &o1)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal(b2, &o2)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(o1, o2), nil
}

//TestResponse ...
func TestResponse(jsonBodyByteArray []byte, fixtureByteArray []byte, t *testing.T) {
	status, err := jsonCompare(fixtureByteArray, jsonBodyByteArray)
	if err == nil {
		assert.Equal(t, status, true)
	} else {
		panic(err.Error())
	}
}

//StartMockServer ...
func StartMockServer(url string, fixtureName string) (func(), string) {
	teardown := testSetup()
	fixture := getFixture(fixtureName)
	mux.HandleFunc(url, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fixture) // nosemgrep : go.lang.security.audit.xss.no-fprintf-to-responsewriter.no-fprintf-to-responsewriter
	})
	return teardown, fixture
}

// Validate webhooks
func VerifyWebhookSignature(requestBody string, webhookSignature string, webhookSecret string) (bool) {
	body := []byte(requestBody)

	isValid := VerifySignature(body, webhookSignature, webhookSecret)

	return isValid
}

func VerifySubscriptionSignature(queryParams map[string]interface{}, webhookSignature string, webhookSecret string) (bool){
     
	payload := fmt.Sprint(queryParams["razorpay_payment_id"], "|", queryParams["razorpay_subscription_id"])

	isValid := VerifySignature([]byte(payload), webhookSignature, webhookSecret)

	return isValid
}

func VerifyPaymentSignature(queryParams map[string]interface{}, webhookSignature string, webhookSecret string) (bool){
     
	payload := fmt.Sprint(queryParams["razorpay_order_id"], "|", queryParams["razorpay_payment_id"])

	isValid := VerifySignature([]byte(payload), webhookSignature, webhookSecret)

	return isValid
}

func VerifyPaymentLinkSignature(queryParams map[string]interface{}, webhookSignature string, webhookSecret string) (bool){
     
	payload := fmt.Sprint(queryParams["payment_link_id"], "|", queryParams["payment_link_reference_id"], "|", 
	queryParams["payment_link_status"], "|", queryParams["razorpay_payment_id"])

	isValid := VerifySignature([]byte(payload), webhookSignature, webhookSecret)

	return isValid
}

func VerifySignature(body []byte, signature string, key string) (bool) {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	if expectedSignature != signature {
		return false
	}
	return true
}
