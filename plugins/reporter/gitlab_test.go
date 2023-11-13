package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/jstang/rasic/types"
)

type MockHTTPClientWithResponse struct{}
type MockHTTPClientWithError struct{}
type MockHTTPClientWithProject struct{}
type MockHTTPClientWithProjects struct{}
type MockHTTPClientWithNextPageHeader struct{}

var mockProject = types.GitlabProject{
	ID:            1,
	WebURL:        "https://gitlab.com/jstang/rasic",
	DefaultBranch: "main",
}

var mockProjects = []types.GitlabProject{
	mockProject,
}

func (m *MockHTTPClientWithResponse) Do(req *http.Request) (*http.Response, error) {
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer([]byte("Test Response"))),
		Status: "200 OK",
	}

	return response, nil
}

func (m *MockHTTPClientWithError) Do(req *http.Request) (*http.Response, error) {
	return nil, errors.New("An Error")
}

func (m *MockHTTPClientWithProject) Do(req *http.Request) (*http.Response, error) {
	project, _ := json.Marshal(&mockProject)
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer(project)),
		Status: "200 OK",
	}

	return response, nil
}

func (m *MockHTTPClientWithProjects) Do(req *http.Request) (*http.Response, error) {
	project, _ := json.Marshal(&mockProjects)
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer(project)),
		Status: "200 OK",
	}

	return response, nil
}

var callCount = 0

func (m *MockHTTPClientWithNextPageHeader) Do(req *http.Request) (*http.Response, error) {
	callCount++

	mockIssues := types.GitlabIssues{
		{
			ID:          callCount,
			Title:       "test-test",
			Description: "test-desc",
			State:       "opened",
		},
	}

	issues, _ := json.Marshal(&mockIssues)

	nextPageHeader := http.Header{}
	if callCount < 5 {
		nextPageHeader = http.Header{"X-Next-Page": []string{strconv.Itoa(callCount)}}
	}
	response := &http.Response{
		Body:   io.NopCloser(bytes.NewBuffer(issues)),
		Header: nextPageHeader,
		Status: "200 OK",
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

func TestGetProject(t *testing.T) {
	httpClient := &MockHTTPClientWithProject{}
	gl := &ReporterGitlab{}
	project := gl.GetProject(httpClient, "jstang/rasic", "1234")
	expected := types.RasicProject{
		ID:            1,
		WebURL:        "https://gitlab.com/jstang/rasic",
		DefaultBranch: "main",
	}
	assert.Equal(t, expected, project)
}

func TestGetProjects(t *testing.T) {
	httpClient := &MockHTTPClientWithProjects{}
	gl := &ReporterGitlab{}
	projects := gl.GetProjects(httpClient, "jstang", "1234")
	expected := []types.RasicProject{
		{
			ID:            1,
			WebURL:        "https://gitlab.com/jstang/rasic",
			DefaultBranch: "main",
		},
	}
	assert.Equal(t, expected, projects)
}

func TestPagination(t *testing.T) {
	httpClient := &MockHTTPClientWithNextPageHeader{}

	var collectedIssues []types.RasicIssue
	collectedIssues = pagination(httpClient, "https://gitlab.com/jstang/rasic", "1234", &collectedIssues, callCount+1)

	assert.Equal(t, 5, len(collectedIssues))
	assert.Equal(t, 1, collectedIssues[0].ID)
	assert.Equal(t, 2, collectedIssues[1].ID)
	assert.Equal(t, 3, collectedIssues[2].ID)
	assert.Equal(t, 4, collectedIssues[3].ID)
	assert.Equal(t, 5, collectedIssues[4].ID)

	issue := collectedIssues[0]
	assert.Equal(t, "test-test", issue.Title)
	assert.Equal(t, "test-desc", issue.Description)
	assert.Equal(t, "opened", issue.State)
}
