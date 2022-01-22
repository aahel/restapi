package config

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	appMiddleware "github.com/aahel/restapi/middleware"
	"github.com/aahel/restapi/server"
	"github.com/go-chi/chi/v5"
	chiMiddleWare "github.com/go-chi/chi/v5/middleware"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"go.elastic.co/ecszap"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Configurations struct {
	DB     DBConfig
	SERVER server.ServerConfig
}

type DBConfig struct {
	URI  string
	NAME string
}

var (
	appConfig *Configurations
	dbConn    *mongo.Database
)

const (
	CovidApiUrl            = "https://data.covid19india.org/v4/min/data.min.json"
	ReverseGeoCodingAPiUrl = "http://api.positionstack.com/v1/reverse"
)

func loadConfig(conf interface{}, file, path string) {
	// Config

	viper.SetConfigName(file) // config file name
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

	err = viper.Unmarshal(conf)
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}

}

func GetAppConfig() *Configurations {
	if appConfig != nil {
		return appConfig
	}
	var conf Configurations
	loadConfig(&conf, "config.yaml", "./")
	appConfig = &conf
	return appConfig
}

func GetConsoleLogger() *zap.SugaredLogger {
	encoder := ecszap.NewDefaultEncoderConfig()

	core := ecszap.NewCore(encoder, os.Stdout, zapcore.DebugLevel)

	logger := zap.New(core, zap.AddCaller())
	return logger.Sugar()
}

func InitRouter(logger *zap.SugaredLogger, appConfig *Configurations) *chi.Mux {
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "OPTIONS", "DELETE"},
		AllowedHeaders: []string{
			"Origin", "Authorization", "Access-Control-Allow-Origin",
			"Access-Control-Allow-Header", "Accept",
			"Content-Type", "X-CSRF-Token",
		},
		ExposedHeaders: []string{
			"Content-Length", "Access-Control-Allow-Origin", "Origin",
		},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(
		cors.Handler,
		chiMiddleWare.RequestID,
		chiMiddleWare.Logger,
		appMiddleware.Recoverer,
	)

	return router
}

func configureDatabase(log *zap.SugaredLogger, conf DBConfig) *mongo.Database {
	clientOptions := options.Client().ApplyURI(conf.URI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client.Database(conf.NAME)
}
func GetDBConn(log *zap.SugaredLogger, conf DBConfig) *mongo.Database {
	if dbConn != nil {
		return dbConn
	}
	dbConn = configureDatabase(log, conf)
	return dbConn
}
