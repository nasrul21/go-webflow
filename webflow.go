package webflow

import (
	"net/http"

	"github.com/nasrul21/go-webflow/client"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/domain"
	"github.com/nasrul21/go-webflow/meta"
	"github.com/nasrul21/go-webflow/site"
)

type Webflow struct {
	Opt        common.Option
	httpClient client.Client
	Meta       meta.Meta
	Domain     domain.Domain
	Site       site.Site
}

func (w *Webflow) init() {
	w.Meta = meta.New(&w.Opt, w.httpClient)
	w.Domain = domain.New(&w.Opt, w.httpClient)
	w.Site = site.New(&w.Opt, w.httpClient)
}

func New(apiKey string) *Webflow {
	webflow := Webflow{
		Opt: common.Option{
			ApiKey:  apiKey,
			BaseURL: "https://api.webflow.com",
		},
		httpClient: &client.ClientImpl{HttpClient: &http.Client{}},
	}

	webflow.init()

	return &webflow
}

func (w *Webflow) WithHttpClient(httpClient client.Client) *Webflow {
	w.httpClient = httpClient
	w.init()
	return w
}
