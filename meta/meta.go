package meta

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nasrul21/go-webflow/client"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/model"
)

type Meta interface {
	GetInfo() (*model.AuthorizationInfo, *common.Error)
	GetInfoWithContext(ctx context.Context) (*model.AuthorizationInfo, *common.Error)
}

type MetaImpl struct {
	Opt        *common.Option
	HttpClient client.HttpClient
}

func New(opt *common.Option, client client.HttpClient) Meta {
	return &MetaImpl{
		Opt:        opt,
		HttpClient: client,
	}
}

func (m *MetaImpl) GetInfo() (*model.AuthorizationInfo, *common.Error) {
	return m.GetInfoWithContext(context.Background())
}

func (m *MetaImpl) GetInfoWithContext(ctx context.Context) (*model.AuthorizationInfo, *common.Error) {
	var response model.AuthorizationInfo
	var header http.Header

	err := m.HttpClient.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/info", m.Opt.BaseURL),
		m.Opt.ApiKey,
		header,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
