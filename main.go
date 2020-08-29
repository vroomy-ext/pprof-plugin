package main

import (
	"net/http/pprof"

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
func Cmdline(ctx vroomy.Context) (res vroomy.Response) {
	pprof.Cmdline(ctx.GetWriter(), ctx.GetRequest())
	return
}

// Profile will serve the pprof.Profile handler
func Profile(ctx vroomy.Context) (res vroomy.Response) {
	pprof.Profile(ctx.GetWriter(), ctx.GetRequest())
	return
}

// Symbol will serve the pprof.Symbol handler
func Symbol(ctx vroomy.Context) (res vroomy.Response) {
	pprof.Symbol(ctx.GetWriter(), ctx.GetRequest())
	return
}

// Trace will serve the pprof.Trace handler
func Trace(ctx vroomy.Context) (res vroomy.Response) {
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
