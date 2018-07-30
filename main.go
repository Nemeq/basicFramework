package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/valyala/fasthttp"
)

func main() {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{})

	rh := InitializeControllers()
	handler := handlerBuilder{}

	handler.Use(LogRequest)
	handler.Use(rh.BuildRouter())

	handler.UsePostRequest(LogError)

	log.Info("Starting server")
	fasthttp.ListenAndServe(":443", handler.Build())
}

func InitializeControllers() *routeHandler {
	rh := NewRouteHanlder()
	rh.AddController(&HomeController{})
	return rh
}

func LogRequest(ctx *fasthttp.RequestCtx) {
	log.Info(fmt.Sprintf("Request type=%v endpoint=%v", string(ctx.Method()), string(ctx.RequestURI())))
}

func LogError(ctx *fasthttp.RequestCtx, errInterface interface{}) {
	statusCode := ctx.Response.StatusCode()
	if statusCode == fasthttp.StatusOK {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		log.Error("Request error ", errInterface)
		ctx.SetBody([]byte(fmt.Sprint(errInterface)))
	}
	if statusCode == fasthttp.StatusNotFound {
		log.Info("Request notfound ", errInterface)
		ctx.Redirect("/NotFound", fasthttp.StatusPermanentRedirect)
	}
}
