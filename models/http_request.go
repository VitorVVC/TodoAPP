package models

import (
	"context"
	"net/http"
)

type HTTPRequest struct {
	Ctx      context.Context
	Method   string
	URL      string
	Body     interface{}
	Response interface{}
	Headers  http.Header
}
