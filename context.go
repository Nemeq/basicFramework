package main

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

type SecContext struct {
	ctx *fasthttp.RequestCtx
}

func (sc *SecContext) Ok(obj interface{}) {
	jsonObj, err := json.Marshal(obj)
	if err != nil {
		panic("Error in Serialization")
	}
	sc.ctx.SetStatusCode(fasthttp.StatusOK)
	sc.ctx.SetBody(jsonObj)
	sc.ctx.SetContentType("application/json")
}

func (sc *SecContext) OkTxt(txt string) {
	sc.response(txt, fasthttp.StatusOK)
}

func (sc *SecContext) NotFoundTxt(txt string) {
	sc.response(txt, fasthttp.StatusNotFound)
}

func (sc *SecContext) response(txt string, statusCode int) {
	sc.ctx.SetBody([]byte(txt))
	sc.ctx.SetStatusCode(statusCode)
}
