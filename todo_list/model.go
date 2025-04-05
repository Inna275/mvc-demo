package main

import (
	"github.com/google/uuid"
)

type Task struct {
	ID        string
	Title     string
	Completed bool
}

type TodoList struct {
	tasks []Task
}

func (t *TodoList) Add(title string) {
	id := uuid.New().String()

	newTask := Task{
		ID:        id,
		Title:     title,
		Completed: false,
	}

	t.tasks = append(t.tasks, newTask)
}

func (t *TodoList) All() []Task {
	return t.tasks
}

func (t *TodoList) Complete(id string) {
	for i, task := range t.tasks {
		if task.ID == id {
			t.tasks[i].Completed = true
			return
		}
	}
}

func (t *TodoList) Remove(id string) {
	for i, task := range t.tasks {
		if task.ID == id {
			t.tasks = append(t.tasks[:i], t.tasks[i+1:]...)
			return
		}
	}
}
