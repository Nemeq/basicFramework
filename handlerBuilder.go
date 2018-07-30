package main

import (
	"github.com/valyala/fasthttp"
)

type HandlerBuilder struct {
	midwares   []func(ctx *fasthttp.RequestCtx)
	errorwares []func(ctx *fasthttp.RequestCtx, errInterface interface{})
}

func (h *HandlerBuilder) Use(midware func(ctx *fasthttp.RequestCtx)) {
	h.midwares = append(h.midwares, midware)
}

func (h *HandlerBuilder) UsePostRequest(midware func(ctx *fasthttp.RequestCtx, errInterface interface{})) {
	h.errorwares = append(h.errorwares, midware)
}

func (h *HandlerBuilder) Build() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		defer func() {
			if r := recover(); r != nil {
				for _, midware := range h.errorwares {
					midware(ctx, r)
				}
			}
		}()
		for _, midware := range h.midwares {
			midware(ctx)
		}
	}
}
