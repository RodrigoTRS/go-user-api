package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"user-api/src/db"
	"user-api/src/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func UpdateUserById(database db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		parsedUUID, uuidParsingError := uuid.Parse(id)
		if uuidParsingError != nil {
			slog.Error("invalid uuid", "error", uuidParsingError)
			utils.SendJSON(
				w,
				utils.Response{Error: "invalid uuid"},
				http.StatusNotFound,
			)
			return
		}

		var receivedUser db.UpdateUserRequest

		if decodingError := json.NewDecoder(r.Body).Decode(&receivedUser); decodingError != nil {
			slog.Error("failed on parsing request body", "error", decodingError)
			utils.SendJSON(
				w,
				utils.Response{Error: "failed on parsing request body"},
				http.StatusInternalServerError,
			)
			return
		}

		user, err := database.Update(parsedUUID, receivedUser)
		if err != nil {
			utils.SendJSON(
				w,
				utils.Response{Error: "user doesn't exists"},
				http.StatusNotFound,
			)
			return
		}

		utils.SendJSON(
			w,
			utils.Response{Data: user},
			http.StatusOK,
		)
	}
}
