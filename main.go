package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Fix: G104
	_ = os.Remove("users.db")

	// Setup database
	db, err := sql.Open("sqlite3", "./users.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create table
	sqlStmt := `CREATE TABLE users (id INTEGER NOT NULL PRIMARY KEY, name TEXT, apikey TEXT);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	// Vulnerable Handler
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")

		// Vulnerable Query
		// query := fmt.Sprintf("SELECT apikey FROM users WHERE id = %s", id)

		// Fix: Parameterized Query
		query := "SELECT apikey FROM users WHERE id = ?"

		// Pass `id` as part of the db.Query
		rows, err := db.Query(query, id)
		if err != nil {
			http.Error(w, "Database error", 500)
			return
		}
		defer rows.Close()

		// Fix: G104
		if _, err := w.Write([]byte("User found (if this were real, a JSON object would be returned)")); err != nil {
			log.Printf("Failed to write response: %v", err)
		}
	})

	fmt.Println("Server running on :8080")

	// Fix: G114
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
