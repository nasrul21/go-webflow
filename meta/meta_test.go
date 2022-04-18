package meta_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/nasrul21/go-webflow"
	"github.com/nasrul21/go-webflow/client/mock"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/model"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	httpClientMockObj := new(mock.ClientMock)
	wf := webflow.New("apikey_123").WithHttpClient(httpClientMockObj)

	httpClientMockObj.CallFunc = func(result interface{}) *common.Error {
		resultString := `{
			"_id": "55818d58616600637b9a5786",
			"createdOn": "2016-10-03T23:12:00.755Z",
			"grantType": "authorization_code",
			"lastUsed": "2016-10-10T21:41:12.736Z",
			"sites": [ ],
			"orgs": [
				"551ad253f0a9c0686f71ed08"
			],
			"workspaces": [ ],
			"users": [
				"545bbecb7bdd6769632504a7"
			],
			"rateLimit": 60,
			"status": "confirmed",
			"application": {
				"_id": "55131cd036c09f7d07883dfc",
				"description": "Testing Application",
				"homepage": "https://webflow.com",
				"name": "Test App",
				"owner": "545bbecb7bdd6769632504a7",
				"ownerType": "Person"
			}
		}`

		_ = json.Unmarshal([]byte(resultString), &result)

		return nil
	}

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
				ID:         "55818d58616600637b9a5786",
				CreatedOn:  time.Date(2016, 10, 03, 23, 12, 00, int(755*time.Millisecond), time.UTC),
				GrantType:  "authorization_code",
				LastUsed:   time.Date(2016, 10, 10, 21, 41, 12, int(736*time.Millisecond), time.UTC),
				Orgs:       []string{"551ad253f0a9c0686f71ed08"},
				Users:      []string{"545bbecb7bdd6769632504a7"},
				Sites:      []interface{}{},
				Workspaces: []interface{}{},
				RateLimit:  60,
				Status:     "confirmed",
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

func TestGetUser(t *testing.T) {
	httpClientMockObj := new(mock.ClientMock)
	wf := webflow.New("apikey_123").WithHttpClient(httpClientMockObj)

	httpClientMockObj.CallFunc = func(result interface{}) *common.Error {
		resultString := `{
			"user": {
				"_id": "545bbecb7bdd6769632504a7",
				"email": "some@email.com",
				"firstName": "Some",
				"lastName": "One"
			}
		}`

		_ = json.Unmarshal([]byte(resultString), &result)

		return nil
	}

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
