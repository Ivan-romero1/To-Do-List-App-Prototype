
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Task struct {

	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var db *pgxpool.Pool

func main() {

	var err error

	databaseURL := os.Getenv("DATABASE_URL")

	for i := 0; i < 30; i++ {

		db, err = pgxpool.New(context.Background(), databaseURL)

		if err == nil {

			err = db.Ping(context.Background())
		}
		if err == nil {

			break
		}
		log.Println("Waiting for PostgreSQL...")

		time.Sleep(2 * time.Second)
	}

	if err != nil {

		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/tasks/", taskHandler)

	log.Println("Go API running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {

	case http.MethodGet:
		rows, err := db.Query(

			r.Context(),
			"SELECT id, description, completed FROM tasks ORDER BY id",
		)
		if err != nil {

			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		tasks := []Task{}

		for rows.Next() {

			var t Task

			if err := rows.Scan(&t.ID, &t.Description, &t.Completed); err != nil {
				
				http.Error(w, err.Error(), 500)
				
				return
			}

			tasks = append(tasks, t)
		}

		if err := rows.Err(); err != nil {

			http.Error(w, err.Error(), 500)

			return
		}

		json.NewEncoder(w).Encode(tasks)

	case http.MethodPost:
		var input struct {

			Description string `json:"description"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			
			http.Error(w, "Invalid JSON", 400)

			return
		}

		input.Description = strings.TrimSpace(input.Description)

		if input.Description == "" {

			http.Error(w, "Description required", 400)

			return
		}

		_, err := db.Exec(
			r.Context(),
			"CALL add_task($1)",
			input.Description,
		)
		if err != nil {

			http.Error(w, err.Error(), 500)

			return
		}

		w.WriteHeader(http.StatusCreated)

	default:

		http.Error(w, "Method not allowed", 405)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/api/tasks/"))

	if err != nil || id <= 0 {

		http.Error(w, "Invalid task ID", 400)
		return
	}

	switch r.Method {

	case http.MethodPatch:

		var input struct {

			Completed bool `json:"completed"`
		}

		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			
			http.Error(w, "Invalid JSON", 400)
			
			return
		}

		result, err := db.Exec(
			r.Context(),
			"UPDATE tasks SET completed=$1 WHERE id=$2",
			input.Completed,
			id,
		)
		if err != nil {

			http.Error(w, err.Error(), 500)

			return
		}
		if result.RowsAffected() == 0 {

			http.Error(w, "Task not found", 404)

			return
		}

		w.WriteHeader(http.StatusNoContent)

	case http.MethodDelete:
		result, err := db.Exec(
			r.Context(),
			"DELETE FROM tasks WHERE id=$1",
			id,
		)
		if err != nil {

			http.Error(w, err.Error(), 500)

			return
		}
		if result.RowsAffected() == 0 {

			http.Error(w, "Task not found", 404)
			
			return
		}

		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", 405)
	}
}
