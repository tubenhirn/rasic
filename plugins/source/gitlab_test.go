package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"gitlab.com/jstang/rasic/types"
)

type MockHTTPClientWithResponse struct{}
type MockHTTPClientWithArray struct{}
type MockHTTPClientWithEmptyObject struct{}
type MockHTTPClientWithError struct{}
type MockHTTPClientWithIssue struct{}

func (m *MockHTTPClientWithResponse) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
		Status: "200 OK",
	}

	return response, nil
}

func (m *MockHTTPClientWithArray) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer([]byte("[]"))),
		Status: "200 OK",
	}

	return response, nil
}

func (m *MockHTTPClientWithEmptyObject) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer([]byte("{}"))),
		Status: "200 OK",
	}

	return response, nil
}

func (m *MockHTTPClientWithError) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("An Error")
}

func (m *MockHTTPClientWithIssue) Do(req *http.Request) (*http.Response, error) {
	responseBody, _ := json.Marshal(&types.GitlabIssue{Title: "test", ID: 666})
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer(responseBody)),
		Status: "201 Created",
	}

	return response, nil
}

func TestApiCallGet(t *testing.T) {
	httpClient := &MockHTTPClientWithResponse{}
	_, err := apiCallGet(httpClient, "", "")
	if err != nil {
		t.Errorf("Shouldn't have received an error with a valid MockHTTPClient, got %s", err)
	}
}

func TestApiCallPost(t *testing.T) {
	httpClient := &MockHTTPClientWithResponse{}
	_, err := apiCallPost(httpClient, "", "", "testbody")
	if err != nil {
		t.Errorf("Shouldn't have received an error with a valid MockHTTPClient, got %s", err)
	}
}

func TestApiCallPostResponse(t *testing.T) {
	httpClient := &MockHTTPClientWithResponse{}
	res, _ := apiCallPost(httpClient, "", "", "testbody")
	if res == nil {
		t.Errorf("Should have received an response with a valid MockHTTPClient, got %s", "nil")
	}
}

func TestApiCallGetResponse(t *testing.T) {
	httpClient := &MockHTTPClientWithResponse{}
	res, _ := apiCallGet(httpClient, "", "")
	if res == nil {
		t.Errorf("Should have received an response with a valid MockHTTPClient, got %s", "nil")
	}
}

func TestApiCallGetWithError(t *testing.T) {
	httpClient := &MockHTTPClientWithError{}
	_, err := apiCallGet(httpClient, "", "")
	t.Log(err)
	if err == nil {
		t.Errorf("Should have received an error with a valid MockHTTPClientWithError, got %s", "nil")
	}
}

func TestApiCallPostWithError(t *testing.T) {
	httpClient := &MockHTTPClientWithError{}
	_, err := apiCallPost(httpClient, "", "", "testbody")
	t.Log(err)
	if err == nil {
		t.Errorf("Should have received an error with a valid MockHTTPClientWithError, got %s", "nil")
	}
}

// func TestGetProjectListWithError(t *testing.T) {
// 	httpClient := &MockHTTPClientWithError{}
// 	_, err := GetProjects(httpClient, "testgroup", "1234")
// 	if err == nil {
// 		t.Errorf("Should have received an error with a valid MockHTTPClientWithError, got %s", "nil")
// 	}
// }

// func TestGetProjectListWithResponse(t *testing.T) {
// 	httpClient := &MockHTTPClientWithArray{}
// 	_, err := GetProjects(httpClient, "testgroup", "1234")
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHTTPClientWithResponse, got %s", err)
// 	}
// }

// func TestGetProjectListWithResponseError(t *testing.T) {
// 	httpClient := &MockHTTPClientWithEmptyObject{}
// 	_, err := GetProjects(httpClient, "testgroup", "1234")
// 	if err == nil {
// 		t.Errorf("Should have received an error with a MockHTTPClientWithEmptyObject, got %s", "nil")
// 	}
// }

// func TestGetProjectWithError(t *testing.T) {
// 	httpClient := &MockHTTPClientWithError{}
// 	_, err := GetProject(httpClient, "testproject", "1234")
// 	if err == nil {
// 		t.Errorf("Should have received an error with a valid MockHTTPClientWithError, got %s", "nil")
// 	}
// }

// func TestGetProjectWithResponse(t *testing.T) {
// 	httpClient := &MockHTTPClientWithEmptyObject{}
// 	_, err := GetProject(httpClient, "testproject", "1234")
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHTTPClientWithResponse, got %s", err)
// 	}
// }

// func TestCreateIssueWithResponse(t *testing.T) {
// 	httpClient := &MockHTTPClientWithIssue{}
// 	issue := &types.RasicIssue{}
// 	_, err := CreateIssue(httpClient, "testproject", "1234", issue)
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHTTPClientWithResponse, got %s", err)
// 	}
// }
