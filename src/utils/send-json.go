package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type Response struct {
	Error string `json:"error,omitempty"`
	Data  any    `json:"data,omitempty"`
}

func SendJSON(w http.ResponseWriter, resp Response, status int) {
	w.Header().Set("Content-Type", "application/json")
	data, marshalErr := json.Marshal(resp)
	if marshalErr != nil {
		slog.Error("failed to marshal json data", "error", marshalErr)
		SendJSON(
			w,
			Response{Error: "something wenr wrong"},
			http.StatusInternalServerError,
		)
		return
	}

	w.WriteHeader(status)
	_, writeError := w.Write(data)
	if writeError != nil {
		slog.Error("failed to write response to client", "error", writeError)
		return
	}
}
