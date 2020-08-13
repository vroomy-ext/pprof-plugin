package main

import (
	"net/http/pprof"

	"github.com/Hatch1fy/httpserve"
	"github.com/vroomy/common"
)

// Init will be called by vroomy on initialization
func Init(env map[string]string) (err error) {
	return
}

// Load will be called by vroomy after plugin initialization
func Load(p common.Plugins) (err error) {
	return
}

// Backend returns the underlying backend to the plugin
func Backend() interface{} {
	return nil
}

// Index will serve the pprof.Index handler
func Index(ctx *httpserve.Context) (res httpserve.Response) {
	pprof.Index(ctx.Writer, ctx.Request)
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
func Trace(ctx *httpserve.Context) (res httpserve.Response) {
	pprof.Trace(ctx.Writer, ctx.Request)
	return
}
