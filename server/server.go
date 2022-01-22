package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

type ServerConfig struct {
	ReadTimeout  int
	WriteTimeout int
	IdleTimeout  int
	ShutdownWait int
	Address      string
}

func StartAndGracefullShutdown(l *zap.SugaredLogger, sm http.Handler, conf ServerConfig) {
	// create a new server
	fmt.Printf("%+v\n", conf)
	s := http.Server{
		Addr:         ":" + conf.Address,                             // configure the bind address
		Handler:      sm,                                             // set the default handler
		ReadTimeout:  time.Duration(conf.ReadTimeout) * time.Second,  // max time to read request from the client
		WriteTimeout: time.Duration(conf.WriteTimeout) * time.Second, // max time to write response to the client
		IdleTimeout:  time.Duration(conf.IdleTimeout) * time.Second,  // max time for connections using TCP Keep-Alive
	}
	// start the server
	go func() {
		l.Infow("Starting server on port " + conf.Address)

		err := s.ListenAndServe()
		if err != nil {
			l.Infof("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)

	// Block until a signal is received.
	sig := <-c
	l.Infow("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		l.Fatal(err)
	}
	error := s.Shutdown(ctx)
	if err != nil {
		l.Fatal(error)
	}

}
