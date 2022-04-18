package domain

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nasrul21/go-webflow/client"
	"github.com/nasrul21/go-webflow/common"
	"github.com/nasrul21/go-webflow/model"
)

type Domain interface {
	GetList(siteID string) ([]model.Domain, *common.Error)
	GetListWithContext(ctx context.Context, siteID string) ([]model.Domain, *common.Error)
}

type DomainImpl struct {
	Opt    *common.Option
	Client client.Client
}

func New(opt *common.Option, client client.Client) Domain {
	return &DomainImpl{
		Opt:    opt,
		Client: client,
	}
}

func (d *DomainImpl) GetList(siteID string) ([]model.Domain, *common.Error) {
	return d.GetListWithContext(context.Background(), siteID)
}
func (d *DomainImpl) GetListWithContext(ctx context.Context, siteID string) ([]model.Domain, *common.Error) {
	response := []model.Domain{}
	var header http.Header

	err := d.Client.Call(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/sites/%s/domains", d.Opt.BaseURL, siteID),
		d.Opt.ApiKey,
		header,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
