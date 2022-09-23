package planet

import (
	"api-sw/src/core/domains/planet/services/create"
	"api-sw/src/shared/providers/httphelper"
	"api-sw/src/shared/tools/communication"
	"net/http"
)

func (handler handler) CreatePlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponet("create_planet")

	handler.Logger.Info(Namespace.Concat("CreatePlanetHandler"), "")

	var dto create.Dto

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Create

	if err := httphelper.GetBody(r.Body, &dto); err != nil {
		handler.Logger.Error(Namespace.Concat("CreatePlanetHandler", "GetBody"), err.Error())

		return comm.ResponseError(400, "error", err)
	}

	return service.Execute(dto)
}
