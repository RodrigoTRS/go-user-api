package routes

import (
	"log/slog"
	"net/http"
	"user-api/src/db"
	"user-api/src/utils"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetUserById(database db.DB) http.HandlerFunc {
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

		user, err := database.FindById(parsedUUID)
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
