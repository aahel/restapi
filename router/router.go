package router

import (
	"context"
	"log"
	"net/http"
	"regexp"

	"github.com/aahel/restapi/errors"
	"github.com/aahel/restapi/respond"
	"github.com/aahel/restapi/types"
)

type Router struct {
	routes []RouteEntry
}

func (rtr *Router) Route(method, path string, handlerFunc http.Handler) {
	exactPath := regexp.MustCompile("^" + path + "$")

	e := RouteEntry{
		Method:      method,
		Path:        exactPath,
		HandlerFunc: handlerFunc,
	}

	rtr.routes = append(rtr.routes, e)
}

func (rtr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("ERROR:", r)
			respond.Fail(w, errors.InternalServerStd())
		}
	}()

	for _, e := range rtr.routes {
		params := e.Match(r)
		if params == nil {
			continue
		}

		ctx := context.WithValue(r.Context(), types.RouteContextKey, params)
		e.HandlerFunc.ServeHTTP(w, r.WithContext(ctx))

		return
	}

	http.NotFound(w, r)
}
