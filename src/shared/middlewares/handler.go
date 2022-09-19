package middlewares

import (
	"api-sw/src/shared/providers/logger"
	"api-sw/src/shared/tools/communication"
	"api-sw/src/shared/tools/namespace"
	"net/http"
)

var NameSpace = namespace.New("shared.middlewares")

func Handler(handler func(r *http.Request) communication.Response) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		NameSpace.AddComponet("handler.handler")

		log := logger.Instance
		response := handler(r)
		if response.Error != nil {
			log.Error(NameSpace.Concat("handler_error"), response.Error.Error())
		}

		if err := response.JSON(w); err != nil {
			log.Error(NameSpace.Concat("response_error"), response.Error.Error())
		}
	}
}
