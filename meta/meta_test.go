package meta_test

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

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

	result.(*model.AuthorizationInfo).ID = "55818d58616600637b9a5786"
	result.(*model.AuthorizationInfo).CreatedOn = time.Date(2016, 10, 10, 23, 12, 00, 00, time.UTC)
	result.(*model.AuthorizationInfo).GrantType = "authorization_code"
	result.(*model.AuthorizationInfo).LastUsed = time.Date(2016, 10, 10, 21, 41, 12, 00, time.UTC)
	result.(*model.AuthorizationInfo).Orgs = []string{"551ad253f0a9c0686f71ed08"}
	result.(*model.AuthorizationInfo).Users = []string{"545bbecb7bdd6769632504a7"}
	result.(*model.AuthorizationInfo).RateLimit = 60
	result.(*model.AuthorizationInfo).Status = "confirmed"
	result.(*model.AuthorizationInfo).Application = model.Application{
		ID:          "55131cd036c09f7d07883dfc",
		Description: "Testing Application",
		Homepage:    "https://webflow.com",
		Name:        "Test App",
		Owner:       "545bbecb7bdd6769632504a7",
		OwnerType:   "Person",
	}

	return nil
}

func TestGetInfo(t *testing.T) {
	httpClientMockObj := new(httpClientMock)
	wf := webflow.New("apikey_123").WithHttpClient(httpClientMockObj)

	testcases := []struct {
		desc        string
		mockClosure func()
		expectedRes *model.AuthorizationInfo
		expectedErr *common.Error
	}{
		{
			desc: "should get info",
			mockClosure: func() {
				httpClientMockObj.On(
					"Call",
					context.Background(),
					http.MethodGet,
					fmt.Sprintf("%s/info", wf.Opt.BaseURL),
					wf.Opt.ApiKey,
					http.Header(nil),
					nil,
					&model.AuthorizationInfo{},
				).Return(nil).Once()
			},
			expectedRes: &model.AuthorizationInfo{
				ID:        "55818d58616600637b9a5786",
				CreatedOn: time.Date(2016, 10, 10, 23, 12, 00, 00, time.UTC),
				GrantType: "authorization_code",
				LastUsed:  time.Date(2016, 10, 10, 21, 41, 12, 00, time.UTC),
				Orgs:      []string{"551ad253f0a9c0686f71ed08"},
				Users:     []string{"545bbecb7bdd6769632504a7"},
				RateLimit: 60,
				Status:    "confirmed",
				Application: model.Application{
					ID:          "55131cd036c09f7d07883dfc",
					Description: "Testing Application",
					Homepage:    "https://webflow.com",
					Name:        "Test App",
					Owner:       "545bbecb7bdd6769632504a7",
					OwnerType:   "Person",
				},
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
					fmt.Sprintf("%s/info", wf.Opt.BaseURL),
					wf.Opt.ApiKey,
					http.Header(nil),
					nil,
					&model.AuthorizationInfo{},
				).Return(common.FromGoErr(fmt.Errorf("some error"))).Once()
			},
			expectedRes: nil,
			expectedErr: common.FromGoErr(fmt.Errorf("some error")),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.mockClosure()

			resp, err := wf.Meta.GetInfo()

			assert.Equal(t, tc.expectedRes, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}

type httpClientMockUser struct {
	mock.Mock
}

func (m *httpClientMockUser) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *common.Error {
	args := m.Called(ctx, method, path, secretKey, header, params, result)
	if args.Get(0) != nil {
		return args.Get(0).(*common.Error)
	}

	result.(*model.AuthorizedUser).User.ID = "545bbecb7bdd6769632504a7"
	result.(*model.AuthorizedUser).User.Email = "some@email.com"
	result.(*model.AuthorizedUser).User.FirstName = "Some"
	result.(*model.AuthorizedUser).User.LastName = "One"

	return nil
}

func TestGetUser(t *testing.T) {
	httpClientMockObj := new(httpClientMockUser)
	wf := webflow.New("apikey_123").WithHttpClient(httpClientMockObj)

	testcases := []struct {
		desc        string
		mockClosure func()
		expectedRes *model.AuthorizedUser
		expectedErr *common.Error
	}{
		{
			desc: "should get user",
			mockClosure: func() {
				httpClientMockObj.On(
					"Call",
					context.Background(),
					http.MethodGet,
					fmt.Sprintf("%s/user", wf.Opt.BaseURL),
					wf.Opt.ApiKey,
					http.Header(nil),
					nil,
					&model.AuthorizedUser{},
				).Return(nil).Once()
			},
			expectedRes: &model.AuthorizedUser{
				User: model.AuthorizedUserDetail{
					ID:        "545bbecb7bdd6769632504a7",
					Email:     "some@email.com",
					FirstName: "Some",
					LastName:  "One",
				},
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
					fmt.Sprintf("%s/user", wf.Opt.BaseURL),
					wf.Opt.ApiKey,
					http.Header(nil),
					nil,
					&model.AuthorizedUser{},
				).Return(common.FromGoErr(fmt.Errorf("some error"))).Once()
			},
			expectedRes: nil,
			expectedErr: common.FromGoErr(fmt.Errorf("some error")),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			tc.mockClosure()

			resp, err := wf.Meta.GetUser()

			assert.Equal(t, tc.expectedRes, resp)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
