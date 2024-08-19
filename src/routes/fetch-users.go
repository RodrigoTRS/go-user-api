package routes

import (
	"net/http"
	"user-api/src/db"
	"user-api/src/utils"
)

func FetchUsers(database db.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users := database.FindAll()

		utils.SendJSON(
			w,
			utils.Response{Data: users},
			http.StatusOK,
		)
	}
}
