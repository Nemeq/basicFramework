package main

import (
	"strings"

	"github.com/valyala/fasthttp"
	)

type routeHandler struct {
	methods map[string](map[string]func(ctx *secContext))
}

const (
	index = "/index"
)

func NewRouteHanlder() *routeHandler {
	routeHandler := routeHandler{
		methods: make(map[string](map[string]func(ctx *secContext))),
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
			method(&secContext{ctx: ctx})
			return
		}
		index, found := rh.methods[GET][uri+index]
		if found {
			index(&secContext{ctx: ctx})
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

func (cr *controllerRouting) Get(route string, execute func(ctx *secContext)) {
	cr.createRoute(GET,route,execute)

}

func (cr *controllerRouting) Post(route string, execute func(ctx *secContext)) {
	cr.createRoute(POST,route,execute)
}

func (cr *controllerRouting) createRoute(method string,route string,execute func(ctx *secContext)) {
	if cr.rh.methods[method] == nil {
		cr.rh.methods[method] = make(map[string]func(ctx *secContext))
	}
	cr.rh.methods[method][strings.ToLower(cr.controller)+"/" + strings.ToLower(route)] = execute
}
