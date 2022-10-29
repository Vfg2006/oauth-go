package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/vfg2006/oauth-go/authenticator"
	"github.com/vfg2006/oauth-go/config"
	"github.com/vfg2006/oauth-go/server/router"

	jsoniter "github.com/json-iterator/go"
	"github.com/sirupsen/logrus"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Server struct {
	httpServer *http.Server
}

func New(
	config *config.Config,
	authenticatorService authenticator.Service,
) (*Server, error) {
	rt := router.New(
		router.WithRoutes(Healthcheck()...),
		router.WithRoutes(AuthenticatorRoutes(authenticatorService)...),
	)

	srv := &Server{
		httpServer: &http.Server{
			Addr:              fmt.Sprintf(":%s", config.Server.Port),
			Handler:           rt,
			ReadHeaderTimeout: 2 * time.Second,
		},
	}

	return srv, nil
}

func (s Server) Run(ctx context.Context) error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			logrus.Error(err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)

	// Block until we receive our signal
	<-signalChan

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return s.Shutdown(ctx)
}

func (s Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
