package mock

import (
	"context"
	"net/http"

	"github.com/nasrul21/go-webflow/common"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
	CallFunc func(result interface{}) *common.Error
}

func (c *ClientMock) Call(ctx context.Context, method string, path string, secretKey string, header http.Header, params interface{}, result interface{}) *common.Error {
	args := c.Called(ctx, method, path, secretKey, header, params, result)
	if args.Get(0) != nil {
		return args.Get(0).(*common.Error)
	}

	return c.CallFunc(&result)
}
