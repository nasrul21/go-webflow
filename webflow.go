package webflow

import (
	"net/http"

	"github.com/nasrul21/go-webflow/client"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/meta"
)

type Webflow struct {
	Opt        common.Option
	httpClient client.HttpClient
	Meta       meta.Meta
}

func (w *Webflow) init() {
	w.Meta = meta.New(&w.Opt, w.httpClient)
}

func New(apiKey string) *Webflow {
	webflow := Webflow{
		Opt: common.Option{
			ApiKey:  apiKey,
			BaseURL: "https://api.webflow.com",
		},
		httpClient: &client.HttpClientImpl{HttpClient: &http.Client{}},
	}

	webflow.init()

	return &webflow
}

func (w *Webflow) WithHttpClient(httpClient client.HttpClient) *Webflow {
	w.httpClient = httpClient
	w.init()
	return w
}
