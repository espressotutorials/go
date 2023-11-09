package vimeo_test

import (
	"encoding/json"
	"fmt"
	"github.com/espressotutorials/go/vimeo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_ListMyProjects(t *testing.T) {
	expect := &vimeo.VimeoResponse[vimeo.Project]{
		Total:  120,
		Paging: vimeo.VimeoPaging{},
		Data:   []vimeo.Project{{Name: "Project 1"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsn, _ := json.Marshal(expect)
		fmt.Fprintf(w, string(jsn))
	}))
	defer srv.Close()

	c := vimeo.Client{
		AccessToken: "ACCESS_TOKEN",
		BaseURL:     srv.URL,
		HttpClient:  &http.Client{},
	}

	res, err := c.ListMyProjects()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expect, res)
}

func TestClient_ListProjectsOfUser(t *testing.T) {
	expect := &vimeo.VimeoResponse[vimeo.Project]{
		Total:  120,
		Paging: vimeo.VimeoPaging{},
		Data:   []vimeo.Project{{Name: "Project 1"}},
	}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsn, _ := json.Marshal(expect)
		fmt.Fprintf(w, string(jsn))
	}))
	defer srv.Close()

	c := vimeo.Client{
		AccessToken: "ACCESS_TOKEN",
		BaseURL:     srv.URL,
		HttpClient:  &http.Client{},
	}

	res, err := c.ListProjectsOfUser(1234)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expect, res)
}

func TestClient_GetMyProject(t *testing.T) {
	expect := &vimeo.Project{URI: "/what/ever/5678", Name: "Project 1"}

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jsn, _ := json.Marshal(expect)
		fmt.Fprintf(w, string(jsn))
	}))
	defer srv.Close()

	c := vimeo.Client{
		AccessToken: "ACCESS_TOKEN",
		BaseURL:     srv.URL,
		HttpClient:  &http.Client{},
	}

	res, err := c.GetMyProject(5648)
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, expect, res)
}
