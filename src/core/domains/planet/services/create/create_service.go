package create

import (
	"api-sw/src/core/domains/planet/repositories"
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"context"
)

type Dto struct {
	Name        string `json:"name" validate:"required"`
	Climate     string `json:"climate" validate:"required"`
	Terrain     string `json:"terrain" validate:"required"`
	Apparitions int    `json:"apparitions"`
}

type Service struct {
	Context    context.Context
	Logger     logger.ILoggerProvider
	Repository repositories.IUserRepository
}

func (service *Service) Execute(Dto) communication.Response {
	comm := communication.New()

	return comm.Response(200, "success_created", "Planeta cadastrado com sucesso")

}
