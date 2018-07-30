package main

type RoutedController interface {
	RegisterRoutes(rh *routeHandler)
}

const (  // Http methods
	GET = "GET"
	POST = "POST"
	PUT = "PUT"
)
