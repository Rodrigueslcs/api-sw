package getbyid

import (
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/tools/namespace"
)

var Namespace = namespace.New("core.domains.planet.get_by_id")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(id string) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()

	dbPlanet, err := service.Repository.FindByID(id)

	if err != nil {
		return comm.ResponseError(400, "planet_not_found", err)
	}
	return comm.Response(200, "success_search", dbPlanet)
}
