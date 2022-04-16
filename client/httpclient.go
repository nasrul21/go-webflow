package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	"github.com/nasrul21/go-webflow/common"
)

type HttpClient interface {
	Call(ctx context.Context, method string, url string, apiKey string, header http.Header, body interface{}, result interface{}) *common.Error
}

type HttpClientImpl struct {
	HttpClient *http.Client
}

func (h *HttpClientImpl) Call(ctx context.Context, method string, url string, apiKey string, header http.Header, body interface{}, result interface{}) *common.Error {
	reqBody := []byte("")
	var req *http.Request
	var err error

	isParamsNil := body == nil || (reflect.ValueOf(body).Kind() == reflect.Ptr && reflect.ValueOf(body).IsNil())

	if !isParamsNil {
		reqBody, err = json.Marshal(body)
		if err != nil {
			return common.FromGoErr(err)
		}
	}

	req, err = http.NewRequestWithContext(
		ctx,
		method,
		url,
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		return common.FromGoErr(err)
	}

	if header != nil {
		req.Header = header
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))
	req.Header.Set("accept-version", "1.0.0")
	req.Header.Set("Content-Type", "application/json")

	return h.doRequest(req, result)
}

func (h *HttpClientImpl) doRequest(req *http.Request, result interface{}) *common.Error {
	resp, err := h.HttpClient.Do(req)
	if err != nil {
		return common.FromGoErr(err)
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.FromGoErr(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return common.FromHTTPErr(resp.StatusCode, respBody)
	}

	if err := json.Unmarshal(respBody, &result); err != nil {
		return common.FromGoErr(err)
	}

	return nil
}
