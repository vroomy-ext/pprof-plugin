package main

import (
	"net/http/pprof"

	"github.com/Hatch1fy/httpserve"
	vroomy "github.com/vroomy/common"
)

// Init will be called by vroomy on initialization
func Init(env map[string]string) (err error) {
	return
}

// Load will be called by vroomy after plugin initialization
func Load(p vroomy.Plugins) (err error) {
	return
}

// Backend returns the underlying backend to the plugin
func Backend() interface{} {
	return nil
}

// Index will serve the pprof.Index handler
func Index(ctx vroomy.Context) (res vroomy.Response) {
	pprof.Index(ctx.GetWriter(), ctx.GetRequest())
	return
}

// Cmdline will serve the pprof.Cmdline handler
func Cmdline(ctx *httpserve.Context) (res httpserve.Response) {
	pprof.Cmdline(ctx.Writer, ctx.Request)
	return
}

// Profile will serve the pprof.Profile handler
func Profile(ctx *httpserve.Context) (res httpserve.Response) {
	pprof.Profile(ctx.Writer, ctx.Request)
	return
}

// Symbol will serve the pprof.Symbol handler
func Symbol(ctx *httpserve.Context) (res httpserve.Response) {
	pprof.Symbol(ctx.Writer, ctx.Request)
	return
}

// Trace will serve the pprof.Trace handler
func Trace(ctx vroomy.Context) (res httpserve.Response) {
	pprof.Trace(ctx.GetWriter(), ctx.GetRequest())
	return
}

// Handler will serve the pprof.Handler handler
func Handler(ctx vroomy.Context) (res vroomy.Response) {
	handlerKey := ctx.Param("handlerKey")
	handler := pprof.Handler(handlerKey)
	handler.ServeHTTP(ctx.GetWriter(), ctx.GetRequest())
	return
}
