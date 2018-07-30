package main

type HomeController struct {
}

func (hc *HomeController) RegisterRoutes(rh *routeHandler) {
	rh.Controller("Home",func (rc *controllerRouting){
		rc.Get("Index", hc.Index)
		rc.Get("GetJson", hc.GetJson)
		rc.Get("GetError", hc.GetError)
	})

}

func (hc *HomeController) Index(sctx *SecContext) {
	sctx.OkTxt("Return message")
}

func (hc *HomeController) GetJson(sctx *SecContext) {
	sctx.Ok("Return message")
}

func (hc *HomeController) GetError(sctx *SecContext) {
	t := 2
	p := 0
	x := t / p
	sctx.Ok(x)
}
