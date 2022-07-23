package utils

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func HandleGetRequest(url string, headers map[string][]string) (*http.Response, error) {
	resp, err := makeRequest("GET", url, nil, headers)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("Error: %d\n%s", resp.StatusCode, b))
	}
	return resp, nil
}

func HandlePostRequest(url string, body []byte, key string, headers map[string][]string) (*http.Response, error) {
	resp, err := makeRequest("POST", url, body, headers)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("Error: %d\n%s", resp.StatusCode, b))
	}
	return resp, nil
}

func HandleHeadRequest(url string, headers map[string][]string) (*http.Response, error) {
	resp, err := makeRequest("HEAD", url, []byte{}, headers)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		b, _ := io.ReadAll(resp.Body)
		return nil, errors.New(fmt.Sprintf("Error: %d\n%s", resp.StatusCode, b))
	}
	return resp, nil
}

func makeRequest(method, url string, b []byte, requestHeaders map[string][]string) (*http.Response, error) {
	headers := map[string][]string{
		"Accept": []string{"application/json"},
	}
	if method == "POST" {
		headers["Content-Type"] = []string{"application/json"}
	}
	for k, v := range requestHeaders {
		headers[k] = v
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header = headers

	return client.Do(req)
}
