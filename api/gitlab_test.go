package api

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

type MockHttpClientWithResponse struct{}
type MockHttpClientWithError struct{}

func (m *MockHttpClientWithResponse) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
	}

	return response, nil
}

func (m *MockHttpClientWithError) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("An Error")
}

func TestApiCall(t *testing.T) {
	httpClient := &MockHttpClientWithResponse{}
	_, err := apiCall(httpClient, "", "")
	if err != nil {
		t.Errorf("Shouldn't have received an error with a valid MockHttpClient, got %s", err)
	}
}

func TestApiCallResponse(t *testing.T) {
	httpClient := &MockHttpClientWithResponse{}
	res, _ := apiCall(httpClient, "", "")
	if res == nil {
		t.Errorf("Should have received an response with a valid MockHttpClient, got %s", "nil")
	}
}

func TestApiCallWithError(t *testing.T) {
	httpClient := &MockHttpClientWithError{}
	_, err := apiCall(httpClient, "", "")
	t.Log(err)
	if err == nil {
		t.Errorf("Should have received an error with a valid MockHttpClientWithError, got %s", "nil")
	}
}
