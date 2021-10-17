package httpserver

import (
	"context"
	"github.com/CodingSquire/bank/pkg/api"
	"github.com/CodingSquire/bank/pkg/httpserver"

	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {
	GetBalance(request api.GetBalanceRequest) (response api.GetBalanceResponse, err error)
}

type getBalanceServer struct {
	transport      GetBalanceTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getBalanceServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetBalance(request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCreateConferenceServer the server creator
func NewGetBalanceServer(transport GetBalanceTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getBalanceServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

// NewPreparedServer factory for server api handler
func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := httpserver.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	getBalanceTransport := NewGetBalanceTransport(httpserver.NewError)

	return httpserver.MakeFastHTTPRouter(
		[]*httpserver.HandlerSettings{
			{
				Path:   URIPathClientGetBalance,
				Method: HTTPMethodGetBalance,
				Handler: NewGetBalanceServer(
					getBalanceTransport,
					svc,
					errorProcessor,
				),
			},
		},
	)
}
