package user

import (
	"chatapp/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Handler struct {
	db *sql.DB
}

func UserHandler(db *sql.DB) *Handler {
	return &Handler{db: db}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &User{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	_, err := user.Create(h.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error creating user: %v", err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user := &User{Username: username}

	err := user.Get(h.db)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Error getting user: %v", err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	user := &User{Username: username}

	if err := user.Delete(h.db); err != nil {
		utils.WriteJSONResponse(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user := &User{Username: username}
	err := user.Get(h.db)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(fmt.Sprintf("Error: %s", err.Error()))
		return
	}

	err = ComparePassword(user.Password, password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"Error": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status": "online",
	})
	return
}

func (h *Handler) handleLogout(w http.ResponseWriter, r *http.Request) {}
