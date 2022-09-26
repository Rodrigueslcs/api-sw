package planet

import (
	"api-sw/src/core/domains/planet/services/update"
	"api-sw/src/shared/providers/httphelper"
	"api-sw/src/shared/tools/communication"
	"fmt"
	"net/http"
)

func (handler handler) UpdatePlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponet("update_handler.UpdatePlanetHandler")

	handler.Logger.Info(Namespace.Component, "")

	var dto update.Dto

	ctx := r.Context()
	id := httphelper.GetParam(r, "id")
	service := handler.Service(ctx).Update
	comm := communication.New()

	if id == "" {
		handler.Logger.Error(Namespace.Concat("IdNotFound"), "ID not found")

		return comm.ResponseError(400, "error_delete", fmt.Errorf("ID not found"))
	}

	if err := httphelper.GetBody(r.Body, &dto); err != nil {
		handler.Logger.Error(Namespace.Concat("UpdatePlanetHandler.GetBody"), err.Error())
		return comm.ResponseError(400, "error_update", err)
	}

	return service.Execute(id, dto)
}
