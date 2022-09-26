package delete

import (
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/tools/namespace"
)

var Namespace = namespace.New("core.domains.planet.delete")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(id string) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()

	err := service.Repository.Delete(id)

	if err != nil {
		return comm.ResponseError(400, "error_delete", err)
	}
	return comm.Response(200, "success", nil)
}
