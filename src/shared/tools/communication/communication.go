package communication

import (
	"encoding/json"
	"net/http"
)

type ResponseError struct {
	Message string `json:"message"`
}

type Response struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   error  `json:"-"`
	Data    any    `json:"data,omitempty"`
}

func (r Response) JSON(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Status)
	if r.Error != nil && r.Data == nil {
		r.Data = r.Error.Error()
	}
	return json.NewEncoder(w).Encode(r)
}
