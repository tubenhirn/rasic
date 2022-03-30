package types

import (
	"net/http"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Issue struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Project struct {
	Id     int
	WebUrl string
}
