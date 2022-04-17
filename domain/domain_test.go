package domain_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/nasrul21/go-webflow"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type httpClientMock struct {
	mock.Mock
}

func (m *httpClientMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *common.Error {
	args := m.Called(ctx, method, path, secretKey, header, params, result)
	if args.Get(0) != nil {
		return args.Get(0).(*common.Error)
	}

	resultString := `[
		{"_id": "589a331aa51e760df7ccb89d","name": "test-api-domain.com"},
		{"_id": "589a331aa51e760df7ccb89e","name": "www.test-api-domain.com"}
	]`

	_ = json.Unmarshal([]byte(resultString), &result)

	return nil
}

func TestGetList(t *testing.T) {
	httpClientMockObj := new(httpClientMock)
	wf := webflow.New("apikey_123").WithHttpClient(httpClientMockObj)

	testcases := []struct {
		desc        string
		mockClosure func()
		expectedRes []model.Domain
		expectedErr *common.Error
	}{
		{
			desc: "should get list domains",
			mockClosure: func() {
				httpClientMockObj.On(
					"Call",
					context.Background(),
					http.MethodGet,
					fmt.Sprintf("%s/sites/%s/domains", wf.Opt.BaseURL, "5ee5e7459c39a7e47341f82f"),
					wf.Opt.ApiKey,
					http.Header(nil),
					nil,
					&[]model.Domain{},
				).Return(nil).Once()
			},
			expectedRes: []model.Domain{
				{ID: "589a331aa51e760df7ccb89d", Name: "test-api-domain.com"},
				{ID: "589a331aa51e760df7ccb89e", Name: "www.test-api-domain.com"},
			},
			expectedErr: nil,
		},
		{
			desc: "should return error",
			mockClosure: func() {
				httpClientMockObj.On(
					"Call",
					context.Background(),
					http.MethodGet,
					fmt.Sprintf("%s/sites/%s/domains", wf.Opt.BaseURL, "5ee5e7459c39a7e47341f82f"),
					wf.Opt.ApiKey,
					http.Header(nil),
					nil,
					&[]model.Domain{},
				).Return(common.FromGoErr(fmt.Errorf("some error"))).Once()
			},
			expectedRes: nil,
			expectedErr: common.FromGoErr(fmt.Errorf("some error")),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.mockClosure()

			resp, err := wf.Domain.GetList("5ee5e7459c39a7e47341f82f")

			assert.Equal(t, tc.expectedRes, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
