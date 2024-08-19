package routes

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"user-api/src/db"
	"user-api/src/utils"
)

func CreateUser(database db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var receivedUser db.CreateUserRequest

		if err := json.NewDecoder(r.Body).Decode(&receivedUser); err != nil {
			slog.Error("failed on parsing request body", "error", err)
			utils.SendJSON(
				w,
				utils.Response{Error: "failed on parsing request body"},
				http.StatusInternalServerError,
			)
			return
		}

		user, err := database.Insert(receivedUser)
		if err != nil {

			utils.SendJSON(
				w,
				utils.Response{Error: "user already exists"},
				http.StatusConflict,
			)
			return
		}

		utils.SendJSON(
			w,
			utils.Response{Data: user},
			http.StatusCreated,
		)
	}
}
