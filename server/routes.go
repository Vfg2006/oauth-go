package server

import (
	"net/http"

	"github.com/vfg2006/oauth-go/authenticator"
	"github.com/vfg2006/oauth-go/server/router"
)

func Healthcheck() []router.Route {
	return []router.Route{
		{
			Path:    "/healthcheck",
			Method:  http.MethodGet,
			Handler: HealthcheckHandler(),
		},
	}
}

func AuthenticatorRoutes(authenticatorService authenticator.Service) []router.Route {
	return []router.Route{
		{
			Path:    "/v1/request",
			Method:  http.MethodPost,
			Handler: RequestToken(authenticatorService),
		},
	}
}
