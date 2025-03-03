package testly

import (
	"bytes"
	"net/http"
)

// Send a GET request to the specified URL
func GetRequest(url string, headers map[string]string, auth *HTTPStepAuth) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if auth.Basic != nil {
		req.SetBasicAuth(auth.Basic.Username, auth.Basic.Password)
	}
	if auth.Bearer != nil {
		req.Header.Set("Authorization", "Bearer "+auth.Bearer.Token)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return client.Do(req)
}

// Send a post request to the specified URL
func PostRequest(url string, headers map[string]string, auth *HTTPStepAuth, body *HTTPStepBody) (*http.Response, error) {
	client := &http.Client{}
	jsonBody, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return nil, err
	}
	if auth.Basic != nil {
		req.SetBasicAuth(auth.Basic.Username, auth.Basic.Password)
	}
	if auth.Bearer != nil {
		req.Header.Set("Authorization", "Bearer "+auth.Bearer.Token)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

// Send a PUT request to the specified URL
func PutRequest(url string, headers map[string]string, auth *HTTPStepAuth, body *HTTPStepBody) (*http.Response, error) {
	client := &http.Client{}
	jsonBody, err := body.MarshalJSON()
	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest("PUT", url, bodyReader)
	if err != nil {
		return nil, err
	}

	if auth.Basic != nil {
		req.SetBasicAuth(auth.Basic.Username, auth.Basic.Password)
	}

	if auth.Bearer != nil {
		req.Header.Set("Authorization", "Bearer "+auth.Bearer.Token)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

// Send a delete request to the specified URL
func DeleteRequest(url string, headers map[string]string, auth *HTTPStepAuth) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	if auth.Basic != nil {
		req.SetBasicAuth(auth.Basic.Username, auth.Basic.Password)
	}
	if auth.Bearer != nil {
		req.Header.Set("Authorization", "Bearer "+auth.Bearer.Token)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}
