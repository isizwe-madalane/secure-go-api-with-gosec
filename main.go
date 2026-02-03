package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Setup a dummy database
	os.Remove("users.db")
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

		query := fmt.Sprintf("SELECT apikey FROM users WHERE id = %s", id)

		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Database error", 500)
			return
		}
		defer rows.Close()

		w.Write([]byte("User found (if this were real, a JSON object would be returned)"))
	})

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
