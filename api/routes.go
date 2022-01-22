package api

import (
	"net/http"

	v1 "github.com/aahel/restapi/api/v1"
	"github.com/go-chi/chi/v5"
)

func InitDocRoutes(r chi.Router, dh http.Handler) {
	r.Method(http.MethodGet, "/docs", dh)
	r.Method(http.MethodGet, "/swagger.yaml", http.FileServer(http.Dir("./")))
}

func InitRecordRoutes(r chi.Router, rh *v1.RecordHandler) {
	r.Method(http.MethodPost, "/v1/records", Handler(rh.GetRecords))
}

func InitInMemoryRoutes(r chi.Router, ih *v1.InMemoryHandler) {
	r.Method(http.MethodGet, "/v1/in-memory", Handler(ih.GetData))
	r.Method(http.MethodPost, "/v1/in-memory", Handler(ih.WriteData))
}
