package create

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

var Namespace = namespace.New("core.domains.planet.services.create.create_service")

type Dto struct {
	Name    string `json:"name" validate:"required"`
	Climate string `json:"climate" validate:"required"`
	Terrain string `json:"terrain" validate:"required"`
	// Apparitions int    `json:"apparitions"`
}

type Service struct {
	Context    context.Context
	Logger     logger.ILoggerProvider
	Repository repositories.IPlanetRepository
}

func (service *Service) Execute(dto Dto) communication.Response {
	service.Logger.Info(Namespace.Concat("Execute"), "")

	comm := communication.New()

	var planet entities.PlanetCreate

	validationErrors := validations.ValidateStruct(dto)

	if len(validationErrors) > 0 {
		service.Logger.Error(Namespace.Concat("Execute", "validate_struct"), fmt.Sprintf("total de erros: %d", len(validationErrors)))
		resp := comm.ResponseError(400, "validate_failed", fmt.Errorf("validate_failed"))
		resp.Data = validationErrors
		return resp
	}

	parse.Unmarshal(dto, &planet)
	planet.Populate()

	dbPlanet, err := service.Repository.FindByName(planet.Name)

	if err != nil {
		return comm.ResponseError(400, "error_create", err)
	}

	if dbPlanet.ID != "" {
		return comm.ResponseError(400, "already_exists", fmt.Errorf("planet already registered "))

	}

	result, err := service.Repository.Create(planet)
	fmt.Println(result)

	if err != nil {
		service.Logger.Error(Namespace.Concat("Execute", "Create"), err.Error())
		return comm.ResponseError(400, "error_create", err)
	}

	return comm.Response(201, "success_create", result)
}
