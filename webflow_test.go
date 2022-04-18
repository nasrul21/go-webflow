package webflow

import (
	"testing"

	"github.com/nasrul21/go-webflow/client"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/meta"
	"github.com/stretchr/testify/assert"
)

func TestWebflowInit(t *testing.T) {
	wf := Webflow{
		Opt: common.Option{
			ApiKey: "apikey_123",
		},
	}
	wf.init()

	assert.Equal(t, *wf.Meta.(*meta.MetaImpl).Opt, wf.Opt)
}

func TestWebflowNew(t *testing.T) {
	wf := New("apikey_123")

	assert.Equal(t, "apikey_123", wf.Opt.ApiKey)
	assert.Equal(t, "https://api.webflow.com", wf.Opt.BaseURL)
}

func TestWebflowWithHttpClient(t *testing.T) {
	httpClient := &client.ClientImpl{}
	wf := New("apikey_123").WithHttpClient(httpClient)

	assert.Equal(t, wf.httpClient, httpClient)
}
