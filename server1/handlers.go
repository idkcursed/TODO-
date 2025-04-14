package main

import (
	"html/template"
	"log"
	"net/http"
	"server/db"
)

type TODO struct {
	ID     int
	Task   string
	Status bool
}

var templ = template.Must(template.ParseFiles("templates/index.html"))

func getalltodos() ([]TODO, error) {
	rows, err := db.DB.Query("SELECT id,task,status FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	todos := []TODO{}
	for rows.Next() {
		var todo TODO
		err = rows.Scan(&todo.ID, &todo.Task, &todo.Status)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
func indexhandler(w http.ResponseWriter, r *http.Request) {
	todos, err := getalltodos()
	if err != nil {
		log.Println(err)
	}
	templ.Execute(w, todos)
}
func addhandler(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("Task")
	_, err := db.DB.Exec("INSERT INTO todos (task) VALUES(?)", task)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
func deletehandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("ID")
	_, err := db.DB.Exec("DELETE FROM todos WHERE id=(?)", id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)

}


func main() {
	// Define routes and their handlers
	http.HandleFunc("/", indexhandler)
	http.HandleFunc("/add", addhandler)
	http.HandleFunc("/delete", deletehandler)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
