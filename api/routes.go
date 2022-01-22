package api

import (
	"net/http"

	v1 "github.com/aahel/restapi/api/v1"
	"github.com/aahel/restapi/router"
	"github.com/aahel/restapi/types"
)

func InitDocRoutes(r *router.Router, dh http.Handler) {
	r.Route(http.MethodGet, "/docs", dh)
	r.Route(http.MethodGet, "/swagger.yaml", http.FileServer(http.Dir("./")))
}

func InitRecordRoutes(r *router.Router, rh *v1.RecordHandler) {
	r.Route(http.MethodPost, "/v1/records", types.Handler(rh.GetRecords))
}

func InitInMemoryRoutes(r *router.Router, ih *v1.InMemoryHandler) {
	r.Route(http.MethodGet, "/v1/in-memory", types.Handler(ih.GetData))
	r.Route(http.MethodPost, "/v1/in-memory", types.Handler(ih.WriteData))
}
