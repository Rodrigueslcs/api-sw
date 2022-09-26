package planet

import (
	"api-sw/src/shared/tools/communication"
	"net/http"
)

func (handler handler) ListPlanetHandler(r *http.Request) communication.Response {
	Namespace.AddComponet("list_handler")
	ctx := r.Context()

	handler.Logger.Info(Namespace.Concat("ListPlanetHandler"), "")

	queryName := r.URL.Query().Get("name")

	if len(queryName) > 0 {
		service := handler.Service(ctx).GetByName
		return service.Execute(queryName)
	}
	queryID := r.URL.Query().Get("id")

	if len(queryID) > 0 {
		service := handler.Service(ctx).GetByID
		return service.Execute(queryID)
	}

	service := handler.Service(ctx).List
	return service.Execute()
}
