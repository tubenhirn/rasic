// package main

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"io/ioutil"
// 	"net/http"
// 	"testing"

// 	"gitlab.com/jstang/rasic/types"
// )

// type MockHttpClientWithResponse struct{}
// type MockHttpClientWithArray struct{}
// type MockHttpClientWithEmptyObject struct{}
// type MockHttpClientWithError struct{}
// type MockHttpClientWithIssue struct{}

// func (m *MockHttpClientWithResponse) Do(req *http.Request) (*http.Response, error) {
// 	response := &http.Response{
// 		Body:   ioutil.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
// 		Status: "200 OK",
// 	}

// 	return response, nil
// }

// func (m *MockHttpClientWithArray) Do(req *http.Request) (*http.Response, error) {
// 	response := &http.Response{
// 		Body:   ioutil.NopCloser(bytes.NewBuffer([]byte("[]"))),
// 		Status: "200 OK",
// 	}

// 	return response, nil
// }

// func (m *MockHttpClientWithEmptyObject) Do(req *http.Request) (*http.Response, error) {
// 	response := &http.Response{
// 		Body:   ioutil.NopCloser(bytes.NewBuffer([]byte("{}"))),
// 		Status: "200 OK",
// 	}

// 	return response, nil
// }

// func (m *MockHttpClientWithError) Do(req *http.Request) (*http.Response, error) {
// 	return nil, errors.New("An Error")
// }

// func (m *MockHttpClientWithIssue) Do(req *http.Request) (*http.Response, error) {
// 	responseBody, _ := json.Marshal(&types.GitlabIssue{Title: "test", ID: 666})
// 	response := &http.Response{
// 		Body:   ioutil.NopCloser(bytes.NewBuffer(responseBody)),
// 		Status: "201 Created",
// 	}

// 	return response, nil
// }

// func TestApiCallGet(t *testing.T) {
// 	httpClient := &MockHttpClientWithResponse{}
// 	_, err := apiCallGet(httpClient, "", "")
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHttpClient, got %s", err)
// 	}
// }

// func TestApiCallPost(t *testing.T) {
// 	httpClient := &MockHttpClientWithResponse{}
// 	_, err := apiCallPost(httpClient, "", "", "testbody")
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHttpClient, got %s", err)
// 	}
// }

// func TestApiCallPostResponse(t *testing.T) {
// 	httpClient := &MockHttpClientWithResponse{}
// 	res, _ := apiCallPost(httpClient, "", "", "testbody")
// 	if res == nil {
// 		t.Errorf("Should have received an response with a valid MockHttpClient, got %s", "nil")
// 	}
// }

// func TestApiCallGetResponse(t *testing.T) {
// 	httpClient := &MockHttpClientWithResponse{}
// 	res, _ := apiCallGet(httpClient, "", "")
// 	if res == nil {
// 		t.Errorf("Should have received an response with a valid MockHttpClient, got %s", "nil")
// 	}
// }

// func TestApiCallGetWithError(t *testing.T) {
// 	httpClient := &MockHttpClientWithError{}
// 	_, err := apiCallGet(httpClient, "", "")
// 	t.Log(err)
// 	if err == nil {
// 		t.Errorf("Should have received an error with a valid MockHttpClientWithError, got %s", "nil")
// 	}
// }

// func TestApiCallPostWithError(t *testing.T) {
// 	httpClient := &MockHttpClientWithError{}
// 	_, err := apiCallPost(httpClient, "", "", "testbody")
// 	t.Log(err)
// 	if err == nil {
// 		t.Errorf("Should have received an error with a valid MockHttpClientWithError, got %s", "nil")
// 	}
// }

// func TestGetProjectListWithError(t *testing.T) {
// 	httpClient := &MockHttpClientWithError{}
// 	_, err := GetProjects(httpClient, "testgroup", "1234")
// 	if err == nil {
// 		t.Errorf("Should have received an error with a valid MockHttpClientWithError, got %s", "nil")
// 	}
// }

// func TestGetProjectListWithResponse(t *testing.T) {
// 	httpClient := &MockHttpClientWithArray{}
// 	_, err := GetProjects(httpClient, "testgroup", "1234")
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHttpClientWithResponse, got %s", err)
// 	}
// }

// func TestGetProjectListWithResponseError(t *testing.T) {
// 	httpClient := &MockHttpClientWithEmptyObject{}
// 	_, err := GetProjects(httpClient, "testgroup", "1234")
// 	if err == nil {
// 		t.Errorf("Should have received an error with a MockHttpClientWithEmptyObject, got %s", "nil")
// 	}
// }

// func TestGetProjectWithError(t *testing.T) {
// 	httpClient := &MockHttpClientWithError{}
// 	_, err := GetProject(httpClient, "testproject", "1234")
// 	if err == nil {
// 		t.Errorf("Should have received an error with a valid MockHttpClientWithError, got %s", "nil")
// 	}
// }

// func TestGetProjectWithResponse(t *testing.T) {
// 	httpClient := &MockHttpClientWithEmptyObject{}
// 	_, err := GetProject(httpClient, "testproject", "1234")
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHttpClientWithResponse, got %s", err)
// 	}
// }

// func TestCreateIssueWithResponse(t *testing.T) {
// 	httpClient := &MockHttpClientWithIssue{}
// 	issue := &types.RasicIssue{}
// 	_, err := CreateIssue(httpClient, "testproject", "1234", issue)
// 	if err != nil {
// 		t.Errorf("Shouldn't have received an error with a valid MockHttpClientWithResponse, got %s", err)
// 	}
// }
