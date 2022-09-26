package planet

import (
	"api-sw/src/shared/providers/httphelper"
	"api-sw/src/shared/tools/communication"
	"net/http"
)

func (handler handler) DeletePlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponet("delete_handler")
	ctx := r.Context()

	handler.Logger.Info(Namespace.Concat("DeletePlanetHandler"), "")

	id := httphelper.GetParam(r, "id")

	service := handler.Service(ctx).Delete
	return service.Execute(id)

}
