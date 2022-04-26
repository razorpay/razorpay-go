package resources

import (
	"fmt"

	"github.com/razorpay/razorpay-go/constants"
	"github.com/razorpay/razorpay-go/requests"
)

//Subscription ...
type Subscription struct {
	Request *requests.Request
}

// All fetches collection of subscription for the given queryParams.
func (s *Subscription) All(queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return s.Request.Get(constants.SUBSCRIPTION_URL, queryParams, extraHeaders)
}

// Fetch fetches a subscription having the given subscriptionID.
func (s *Subscription) Fetch(subscriptionID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Get(url, queryParams, extraHeaders)
}

// Create creates a new subscription for the given data.
func (s *Subscription) Create(data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	return s.Request.Post(constants.SUBSCRIPTION_URL, data, extraHeaders)
}

// Cancel cancels a subscription having the given subscriptionID.
func (s *Subscription) Cancel(subscriptionID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Post(url, data, extraHeaders)
}

// CreateAddon creates a new addon on the subscription having the given subscriptionID.
func (s *Subscription) CreateAddon(subscriptionID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/addons", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Post(url, data, extraHeaders)
}

// Pause a subscription having the given subscriptionID.
func (s *Subscription) Pause(subscriptionID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/pause", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Post(url, data, extraHeaders)
}

// Resume resumes a subscription having the given subscriptionID.
func (s *Subscription) Resume(subscriptionID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/resume", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Post(url, data, extraHeaders)
}

// PendingUpdate fetches details of a pending update.
func (s *Subscription) PendingUpdate(subscriptionID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/retrieve_scheduled_changes", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Get(url, queryParams, extraHeaders)
}

// CancelScheduledChanges cancels a pending update.
func (s *Subscription) CancelScheduledChanges(subscriptionID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/cancel_scheduled_changes", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Post(url, data, extraHeaders)
}

// Update updates a subscription for the given data.
func (s *Subscription) Update(subscriptionID string, data map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s", constants.SUBSCRIPTION_URL, subscriptionID)
	return s.Request.Patch(url, data, extraHeaders)
}

// DeleteOffer deletes a offer linked to subscription.
func (s *Subscription) DeleteOffer(subscriptionID string, offerID string, queryParams map[string]interface{}, extraHeaders map[string]string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/%s/%s", constants.SUBSCRIPTION_URL, subscriptionID, offerID)
	return s.Request.Delete(url, queryParams, extraHeaders)
}
