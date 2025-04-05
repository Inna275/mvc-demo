package main

import (
	"net/http"
	"strings"
)

var todoList = &TodoList{}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, todoList.All())
}

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		todoList.Add(title)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleRemoveTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/remove/")
	todoList.Remove(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func HandleCompleteTask(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/complete/")
	todoList.Complete(id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
