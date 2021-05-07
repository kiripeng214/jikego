package xnet

import (
	"context"
	"net/http"
	"os"
)

type Option func(o *options)

//继续封装
type options struct {
	HttpServer *http.Server
	ctx        context.Context
	signals    []os.Signal
}

func HttpSever(srv *http.Server) Option {
	return func(o *options) {
		o.HttpServer = srv
	}
}
