package planet

import (
	"api-sw/src/core/domains/planet/services/create"
	"api-sw/src/shared/providers/httphelper"
	"api-sw/src/shared/tools/communication"
	"net/http"
)

func (handler handler) CreateHandler(r *http.Request) communication.Response {
	var dto create.Dto

	ctx := r.Context()
	comm := communication.New()
	service := handler.Service(ctx).Create

	if err := httphelper.GetBody(r.Body, &dto); err != nil {
		return comm.ResponseError(400, "error", err)
	}

	return service.Execute(dto)
}
