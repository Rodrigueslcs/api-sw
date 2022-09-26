package update

import (
	"api-sw/src/core/domains/planet/entities"
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/tools/namespace"
	"api-sw/src/shared/tools/parse"
	"api-sw/src/shared/validations"
	"context"
	"fmt"
)

var Namespace = namespace.New("core.domains.planet.services.update.update_service")

type Dto struct {
	Name    string `json:"name" validate:"required"`
	Climate string `json:"climate" validate:"required"`
	Terrain string `json:"terrain" validate:"required"`
}

type Service struct {
	Context    context.Context
	Repository repositories.IPlanetRepository
	Logger     logger.ILoggerProvider
}

func (service *Service) Execute(id string, dto Dto) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	var planet entities.PlanetUpdate

	comm := communication.New()

	validationErrors := validations.ValidateStruct(dto)
	if len(validationErrors) > 0 {
		service.Logger.Error(Namespace.Concat("Execute,ValidateStruct"), fmt.Sprintf("total de erros %d", len(validationErrors)))
		response := comm.ResponseError(400, "validate_failed", fmt.Errorf("validation errors"))
		response.Data = validationErrors
		return response
	}
	dbPlanet, err := service.Repository.FindByName(planet.Name)

	if err != nil {
		return comm.ResponseError(400, "error_update", err)
	}

	if dbPlanet.ID != "" && dbPlanet.ID != id {
		return comm.ResponseError(400, "already_exist", fmt.Errorf("planet exitent"))
	}

	parse.Unmarshal(dto, &planet)

	result, err := service.Repository.Update(id, planet)
	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "Update"), err.Error())
		return comm.ResponseError(400, "error_update", err)
	}

	return comm.Response(200, "success_update", result)
}
