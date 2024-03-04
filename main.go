package main

import (
	"net/http"
	"strconv"

	"todo/internal/models"
	"todo/internal/storage"
	"todo/internal/todos"
	"todo/view/layout"
	todoview "todo/view/todo"

	"github.com/a-h/templ"
)

func main() {

	todosDB := storage.NewDumpStorage()

	todoService := todos.NewService(todosDB)

	http.Handle("/", templ.Handler(layout.Base()))

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		list, _ := todoService.List()
		todoview.TodoList(list).Render(r.Context(), w)
	})

	http.HandleFunc("/todos/add", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		c := r.FormValue("content")

		t, _ := todoService.Create(c)
		todoview.Entry(*t).Render(r.Context(), w)
	})

	http.HandleFunc("/todos/edit", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		t, _ := todoService.Find(int64(id))
		todoview.Edit(*t).Render(r.Context(), w)
	})

	http.HandleFunc("/todos/update", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		r.ParseForm()
		content := r.FormValue("content")
		t := models.Instance{Id: int64(id), Content: content, Completed: false}
		updatedTodo, _ := todoService.Update(&t)
		todoview.Entry(*updatedTodo).Render(r.Context(), w)
	})

	http.HandleFunc("/todos/delete", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		todoService.Delete(int64(id))
	})

	http.HandleFunc("/todos/search", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		term := r.FormValue("match")
		list, _ := todoService.Search(term)
		todoview.TodoList(list).Render(r.Context(), w)
	})

	http.HandleFunc("/todos/completed", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		todoService.SetCompleted(int64(id), true)
		t, _ := todoService.Find(int64(id))
		todoview.Entry(*t).Render(r.Context(), w)
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		layout.Base().Render(r.Context(), w)
	})

	http.ListenAndServe(":80", nil)
}
