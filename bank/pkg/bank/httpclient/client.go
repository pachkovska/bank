package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	"github.com/CodingSquire/bank/pkg/api"
)

// Service implements Service interface
type Service interface {
	GetBalance(request *api.GetBalanceRequest) (response api.GetBalanceResponse, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetBalance GetBalanceTransport
}

// GetBrandsByID ...
func (s *client) GetBalance(request *api.GetBalanceRequest) (response api.GetBalanceResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	ctx := context.Background()
	if err = s.transportGetBalance.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetBalance.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,
	transportGetBalance GetBalanceTransport,
) Service {
	return &client{
		cli: cli,

		transportGetBalance: transportGetBalance,
	}
}

//// NewPreparedClient create and set up http client
//func NewPreparedClient(
//	serverURL string,
//	serverHost string,
//	maxConns int,
//	errorProcessor errorProcessor,
//	errorCreator errorCreator,
//
//	uriPathCreateThesis string,
//
//	httpMethodGetBrandsByID string,
//
//) Service {
//
//	transportCreateThesis := NewCreateThesisTransport(
//		errorProcessor,
//		errorCreator,
//		serverURL+uriPathCreateThesis,
//		httpMethodGetBrandsByID,
//	)
//
//	return NewClient(
//		&fasthttp.HostClient{
//			Addr:     serverHost,
//			MaxConns: maxConns,
//		},
//
//		transportCreateThesis,
//	)
//}
