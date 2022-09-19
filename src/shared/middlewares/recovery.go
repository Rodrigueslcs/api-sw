package middlewares

import (
	"api-sw/src/shared/tools/communication"
	"errors"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				comm := communication.New()
				// span, ok := trancer.SpanFromContext(r.Context())
				response := comm.ResponseError(500, "error", errors.New("internal server error"))
				response.JSON(w)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
