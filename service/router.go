package service

import "net/http"

// Contains all needed information for a http route.
type Route struct {
	// Path component of url.
	path string
	// Handler function working on the requests on this route.
	handler http.HandlerFunc
	// HTTP-Method, this route is for.
	method string
}

// Contains all configured routes of a router.
type Router struct {
	routes []Route
}

// Creates a new Router with initialized routes slice.
func NewRouter() *Router {
	return &Router{routes: []Route{}}
}

// Adds a route with method, path and handler func to rt.
func (rt *Router) AddRoute(method string, path string, handler http.HandlerFunc) {
	rt.routes = append(rt.routes, Route{method: method, path: path, handler: handler})
}

// Implements the ServeHTTP method of the http.Handler interface for the Router type.
func (rt *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedPath := r.URL.Path
	for _, route := range rt.routes {
		if route.path == requestedPath && route.method == r.Method {
			route.handler(w, r)
			return
		}
	}
	http.Error(w, "resource not found", http.StatusNotFound)
}
