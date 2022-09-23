package planet

import (
	"api-sw/src/shared/tools/communication"
	"net/http"
)

func (handler handler) ListPlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponet("list_handler")

	handler.Logger.Info(Namespace.Concat("ListPlanetHandler"), "")

	ctx := r.Context()
	service := handler.Service(ctx).List
	return service.Execute()
}
