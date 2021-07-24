package resources

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
)

// This function takes in the raw request body and returns a byte slice to be used in json unmarshalling
/* Usage

body, err := client.resources.VerifyWebhookSignature(r.Body, webhookSignature, webhookSecret)
if err != nil {
	 # Handle error
}

# Proceed with further processing. Eg: JSON unmarshalling

json.Unmarshal(body, &variable)

*/
func VerifyWebhookSignature(requestBody io.ReadCloser, webhookSignature string, webhookSecret string) ([]byte, error) {
	body, err := ioutil.ReadAll(requestBody)
	if err != nil {
		return nil, err
	}
	if err := VerifySignature(body, webhookSignature, webhookSecret); err != nil {
		return nil, err
	}
	return body, nil
}

func VerifySignature(body []byte, signature string, key string) error {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	if expectedSignature != signature {
		return errors.New("Signatures don't match")
	}
	return nil
}
