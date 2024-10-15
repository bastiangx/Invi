package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type RSVP struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Event string `json:"event"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	var rsvp RSVP
	err := json.NewDecoder(r.Body).Decode(&rsvp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	dbPath := os.Getenv("LAMBDA_TASK_ROOT") + "/../db/rsvp.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		http.Error(w, "Database connection failed", http.StatusInternalServerError)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO rsvp(name, email, event) VALUES (?, ?, ?)")
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(rsvp.Name, rsvp.Email, rsvp.Event)
	if err != nil {
		http.Error(w, "Failed to execute statement", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "RSVP added for %s", rsvp.Name)
}

func main() {
	http.HandleFunc("/", handler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
