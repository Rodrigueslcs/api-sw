package list

import (
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/tools/namespace"

	"go.mongodb.org/mongo-driver/bson"
)

var Namespace = namespace.New("core.domains.planet.services.list.list_service")

type Service struct {
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute() communication.Response {
	service.Logger.Info(Namespace.Concat("Excute"), "")

	comm := communication.New()

	filter := bson.M{}
	documents, err := service.Repository.FindAll(filter)

	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "FindAll"), err.Error())
		return comm.ResponseError(400, "error_list", err)
	}

	return comm.Response(200, "success", documents)
}
