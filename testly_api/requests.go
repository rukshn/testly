package testly

import (
	"net/http"
)

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
