package router

import (
	"net/http"
	"regexp"
	"strings"
)

type RouteEntry struct {
	Path        *regexp.Regexp
	Method      string
	HandlerFunc http.Handler
}

func (ent *RouteEntry) Match(r *http.Request) map[string]string {
	match := ent.Path.FindStringSubmatch(r.URL.Path)
	if match == nil || r.Method != ent.Method {
		return nil
	}

	params := make(map[string]string)
	for i, group := range r.URL.Query() {
		params[strings.ToLower(i)] = strings.Join(group, ",")
	}

	return params
}
