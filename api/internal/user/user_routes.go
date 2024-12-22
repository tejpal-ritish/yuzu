package user

import (
	"chatapp/utils"
	"net/http"
)

func UserRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			h.CreateUser(w, r)
		case http.MethodGet:
			h.GetUser(w, r)
		case http.MethodDelete:
			h.DeleteUser(w, r)
		default:
			utils.WriteJSONResponse(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		}
	})
	mux.HandleFunc("POST /login", h.handleLogin)
}
