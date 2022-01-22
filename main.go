package main

import (
	"github.com/aahel/restapi/api"
	handler "github.com/aahel/restapi/api/v1"
	"github.com/aahel/restapi/config"
	"github.com/aahel/restapi/server"
	"github.com/aahel/restapi/service"
	"github.com/aahel/restapi/store"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	lgr := config.GetConsoleLogger()
	appConfig := config.GetAppConfig()
	dbConn := config.GetDBConn(lgr, appConfig.DB)
	inMemory := map[string]string{}
	router := config.InitRouter(lgr, appConfig)

	recordsStore := store.NewRecordsStore(lgr, dbConn)
	recordsService := service.NewRecordService(lgr, recordsStore)
	recordsHandler := handler.NewRecordHandler(lgr, recordsService)
	api.InitRecordRoutes(router, recordsHandler)

	inMemoryStore := store.NewInMemoryStore(lgr, inMemory)
	inMemoryService := service.NewInMemoryService(lgr, inMemoryStore)
	inMemoryHandler := handler.NewInMemoryHandler(lgr, inMemoryService)
	api.InitInMemoryRoutes(router, inMemoryHandler)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	dh := middleware.Redoc(opts, nil)
	api.InitDocRoutes(router, dh)

	server.StartAndGracefullShutdown(lgr, router, appConfig.SERVER)
}
