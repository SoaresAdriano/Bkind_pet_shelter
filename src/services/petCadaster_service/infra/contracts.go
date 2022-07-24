package infra

import (
	"context"

	"github.com/ditointernet/go-dito/lib/http"
)

type HttpClientGet interface {
	Get(ctx context.Context, request http.HTTPRequest) (http.HTTPResult, error)
}

type HttpClientPost interface {
	Post(ctx context.Context, request http.HTTPRequest) (http.HTTPResult, error)
}

type HttpClient interface {
	HttpClientGet
	HttpClientPost
}
