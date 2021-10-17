package httpclient

import (
	"context"
	"encoding/json"
	"github.com/CodingSquire/bank/pkg/api"
	"net/http"

	"github.com/valyala/fasthttp"
)

type errorCreator func(status int, format string, v ...interface{}) error

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}

// GetBalanceTransport transport interface
type GetBalanceTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *api.GetBalanceRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response api.GetBalanceResponse, err error)
}

type getBalanceTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *getBalanceTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *api.GetBalanceRequest) (err error) {
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(request)
	if err != nil {
		return
	}
	if _, err = fasthttp.WriteBrotli(r.BodyWriter(), body); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// DecodeResponse method for decoding response on client side
func (t *getBalanceTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response api.GetBalanceResponse, err error) {
	err = json.Unmarshal(r.Body(), &response)
	if err != nil {
		return response, t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// NewGetBalanceTransport the transport creator for http requests
func NewGetBalanceTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) GetBalanceTransport {
	return &getBalanceTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}
