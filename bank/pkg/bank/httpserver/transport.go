package httpserver

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/valyala/fasthttp"

	"github.com/CodingSquire/bank/pkg/api"
)

type errorCreator func(status int, format string, v ...interface{}) error

// swagger:route POST /conference/create conference-tag CreateConferenceRequest
// CreateConferenceTransport does Create Conference.
// responses:
//   200: CreateConferenceResponse

// GetBalanceTransport transport interface
type GetBalanceTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetBalanceRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetBalanceResponse) (err error)
}

type getBalanceTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getBalanceTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetBalanceRequest, err error) {
	err = json.Unmarshal(r.Body(), &request)
	if err != nil {
		return request, t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getBalanceTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetBalanceResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	body, err := json.Marshal(response)
	if err != nil {
		return
	}
	if _, err = fasthttp.WriteBrotli(r.BodyWriter(), body); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return

}

// NewGetBalanceTransport the transport creator for http requests
func NewGetBalanceTransport(errorCreator errorCreator) GetBalanceTransport {
	return &getBalanceTransport{
		errorCreator: errorCreator,
	}
}
