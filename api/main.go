package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"chatapp/db"
	"chatapp/internal/user"
)

type Server struct {
	db *sql.DB
}

func main() {
	// Establish connection to db
	db, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Error connecting to db:", err)
	}
	defer db.Close()

	s := &Server{db: db}
	mux := http.NewServeMux()

	h := user.UserHandler(s.db)
	user.UserRoutes(mux, h)
	// mux.Handle("/ws", websocket.Handler(handleWS))

	http.Handle("/", enableCors(mux))
	http.ListenAndServe(":3000", nil)
}

func enableCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
