package main

import (
	"strings"

	"github.com/valyala/fasthttp"
	)

type routeHandler struct {
	methods map[string](map[string]func(ctx *SecContext))
}

const (
	index = "/index"
)

func NewRouteHanlder() *routeHandler {
	routeHandler := routeHandler{
		methods: make(map[string](map[string]func(ctx *SecContext))),
	}
	routeHandler.AddController(&NotFoundController{})
	return &routeHandler
}

func (rh *routeHandler) AddController(controller RoutedController) {
	controller.RegisterRoutes(rh)
}

func (rh *routeHandler) BuildRouter() func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		uri := strings.ToLower(string(ctx.RequestURI()))[1:]

		method, found := rh.methods[string(ctx.Method())][uri]
		if found {
			method(&SecContext{ctx: ctx})
			return
		}
		index, found := rh.methods[GET][uri+index]
		if found {
			index(&SecContext{ctx: ctx})
			return
		}

		NotFound(ctx)
	}
}

func (rh *routeHandler) Controller(cname string,cr func(rh *controllerRouting)) {
	cr(&controllerRouting{rh: rh,controller:cname })
}

func NotFound(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)
	panic("404 Not found")
}

type controllerRouting struct {
	controller string
	rh *routeHandler
}

func (cr *controllerRouting) Get(route string, execute func(ctx *SecContext)) {
	cr.createRoute(GET,route,execute)

}

func (cr *controllerRouting) Post(route string, execute func(ctx *SecContext)) {
	cr.createRoute(POST,route,execute)
}

func (cr *controllerRouting) createRoute(method string,route string,execute func(ctx *SecContext)) {
	if cr.rh.methods[method] == nil {
		cr.rh.methods[method] = make(map[string]func(ctx *SecContext))
	}
	cr.rh.methods[method][strings.ToLower(cr.controller)+"/" + strings.ToLower(route)] = execute
}
