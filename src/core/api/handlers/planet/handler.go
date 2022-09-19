package planet

import (
	"api-sw/src/core/domains/planet/services"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/tools/namespace"
	"context"
	"net/http"
)

var Namespace = namespace.New("core.api.handlers.planet")

type IHandler interface {
	CreateHandler(r *http.Request) communication.Response
}

type handler struct {
	Logger logger.ILoggerProvider
}

func NewHandler(logger logger.ILoggerProvider) handler {
	return handler{
		Logger: logger,
	}
}

func (h handler) Service(ctx context.Context) *services.Services {

	dependencies := services.Dependecies{
		Context:    ctx,
		Repository: nil,
		Logger:     logger.New(),
	}

	return services.NewUser(dependencies)
}
