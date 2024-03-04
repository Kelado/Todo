package main

import (
	"html/template"
	"net/http"
	"strconv"

	"todo/internal/models"
	"todo/internal/storage"
	"todo/internal/todos"
)

func main() {
	tmpl := make(map[string]*template.Template)
	tmpl["todo"] = template.Must(template.ParseFiles("templates/todo-block.html", "templates/todo.html"))
	tmpl["todos"] = template.Must(template.ParseFiles("templates/todos.html", "templates/todo.html"))
	tmpl["todo-edit"] = template.Must(template.ParseFiles("templates/todo-edit.html"))

	todosDB := storage.NewDumpStorage()

	todoService := todos.NewService(todosDB)

	http.Handle("/", http.FileServer(http.Dir("public/")))

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		list, _ := todoService.List()
		tmpl["todos"].Execute(w, list)
	})

	http.HandleFunc("/todos/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		c := r.FormValue("content")

		t, _ := todoService.Create(c)
		tmpl["todo"].Execute(w, t)
	})

	http.HandleFunc("/todos/edit", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		t, _ := todoService.Find(int64(id))
		tmpl["todo-edit"].Execute(w, t)
	})

	http.HandleFunc("/todos/update", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		r.ParseForm()
		content := r.FormValue("content")
		t := models.Instance{Id: int64(id), Content: content, Completed: false}
		updatedTodo, _ := todoService.Update(&t)
		tmpl["todo"].Execute(w, updatedTodo)
	})

	http.HandleFunc("/todos/delete", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		todoService.Delete(int64(id))
	})

	http.HandleFunc("/todos/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		term := r.FormValue("match")
		list, _ := todoService.Search(term)
		tmpl["todos"].Execute(w, list)
	})

	http.HandleFunc("/todos/completed", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		todoService.SetCompleted(int64(id), true)
		t, _ := todoService.Find(int64(id))
		tmpl["todo"].Execute(w, t)
	})

	http.ListenAndServe(":80", nil)
}
