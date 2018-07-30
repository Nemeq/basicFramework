package main

type NotFoundController struct {
}

func (hc *NotFoundController) RegisterRoutes(rh *routeHandler) {
	rh.Controller("NotFound",func (rc *controllerRouting){
		rc.Get("Index",hc.Index)
	})
}

func (hc *NotFoundController) Index(sctx *SecContext) {
	sctx.NotFoundTxt("Page not found")
}

