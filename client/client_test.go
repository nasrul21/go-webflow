package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nasrul21/go-webflow/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ClientTestSuite struct {
	suite.Suite
	client client.Client
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}

func (c *ClientTestSuite) SetupSuite() {
	c.client = &client.ClientImpl{
		HttpClient: &http.Client{},
	}
}

func (c *ClientTestSuite) TestCallSuccess() {
	json := `{"first_name": "some", "last_name": "one"}`
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(json))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		http.MethodGet,
		server.URL,
		"apikey_123",
		http.Header{},
		map[string]interface{}{"hello": "world"},
		&result,
	)

	expectedRes := map[string]interface{}{
		"first_name": "some",
		"last_name":  "one",
	}

	assert.Nil(c.T(), err)
	assert.Equal(c.T(), expectedRes, result)
}

func (c *ClientTestSuite) TestCallErrorDecodeResponse() {
	json := ``
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(json))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		http.MethodGet,
		server.URL,
		"apikey_123",
		http.Header{},
		map[string]interface{}{"hello": "world"},
		&result,
	)

	expectedRes := map[string]interface{}{}

	assert.NotNil(c.T(), err)
	assert.Equal(c.T(), expectedRes, result)
}

func (c *ClientTestSuite) TestCallErrorStatusCode() {
	json := `{"err":"something went wrong!"}`
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(json))
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		http.MethodGet,
		server.URL,
		"apikey_123",
		http.Header{},
		map[string]interface{}{"hello": "world"},
		&result,
	)

	expectedRes := map[string]interface{}{}

	assert.Equal(c.T(), "something went wrong!", err.Err)
	assert.Equal(c.T(), expectedRes, result)
}

func (c *ClientTestSuite) TestCallErrorDo() {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	server.URL = ""
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		http.MethodGet,
		server.URL,
		"apikey_123",
		http.Header{},
		map[string]interface{}{"hello": "world"},
		&result,
	)

	assert.Len(c.T(), result, 0)
	assert.Contains(c.T(), err.Message, "unsupported protocol scheme")
}

func (c *ClientTestSuite) TestCallErrorRequestBody() {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		http.MethodGet,
		server.URL,
		"apikey_123",
		http.Header{},
		func() {},
		&result,
	)

	assert.Len(c.T(), result, 0)
	assert.Contains(c.T(), err.Message, "unsupported type")
}

func (c *ClientTestSuite) TestCallErrorRequest() {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		"wrong method",
		server.URL,
		"apikey_123",
		http.Header{},
		map[string]interface{}{"hello": "world"},
		&result,
	)

	assert.Len(c.T(), result, 0)
	assert.Contains(c.T(), err.Message, "invalid method")
}

func (c *ClientTestSuite) TestCallErrorResponseBody() {
	dummyHandler := func(w http.ResponseWriter, r *http.Request) {
		// w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Length", "1")
	}

	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	result := map[string]interface{}{}

	err := c.client.Call(
		context.Background(),
		http.MethodGet,
		server.URL,
		"apikey_123",
		http.Header{},
		map[string]interface{}{"hello": "world"},
		&result,
	)

	assert.Len(c.T(), result, 0)
	assert.Contains(c.T(), err.Message, "unexpected EOF")
}
