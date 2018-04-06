package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	razorpay "github.com/razorpay/razorpay-go"
)

var (
	//Mux ...
	Mux *http.ServeMux
	//Server ...
	Server *httptest.Server
	//Client ...
	Client *razorpay.Client
)

const (
	//TestAccessKey ...
	TestAccessKey = "fake_key"
	//TestAccessSecret ...
	TestAccessSecret = "fake_secret"
)

//TestSetup ...
func TestSetup() func() {
	Mux = http.NewServeMux()
	Server = httptest.NewServer(Mux)
	Client = razorpay.NewClient(TestAccessKey, TestAccessSecret)
	razorpay.Request.BaseURL = Server.URL

	return func() {
		Server.Close()
	}
}

//GetFixture ...
func GetFixture(path string) string {
	filepath := path + ".json"
	b, err := ioutil.ReadFile("../testdata/" + filepath)
	if err != nil {
		panic(err)
	}
	// Remove newlines and extra spaces
	return strings.Replace(strings.Replace(string(b), "\n", "", -1), " ", "", -1)
}

//JSONCompare ...
func JSONCompare(s1, s2 string) (bool, error) {
	var o1, o2 interface{}
	err := json.Unmarshal([]byte(s1), &o1)
	if err != nil {
		return false, err
	}
	err = json.Unmarshal([]byte(s2), &o2)
	if err != nil {
		return false, err
	}
	return reflect.DeepEqual(o1, o2), nil
}
