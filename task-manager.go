package main

import "fmt"

type Task struct {
	Title       string
	Description string
	DueDate     string
	Completed   bool
}

func AddTask(tasks *[]Task, task Task) {
	*tasks = append(*tasks, task)
}

func RemoveTask(tasks *[]Task, title string) {
	for i := range *tasks {
		if (*tasks)[i].Title == title {
			*tasks = append((*tasks)[:i], (*tasks)[i+1:]...)
			fmt.Printf("Removed Task: %v\n", title)
			return
		}
	}
	fmt.Printf("Task not found with title:%s\n", title)
}

func CompleteTask(tasks *[]Task, title string) {
	for i := range *tasks {
		if (*tasks)[i].Title == title {
			(*tasks)[i].Completed = true
		}
	}
	fmt.Printf("Task not found with title:%s\n", title)
}
func ListTasks(tasks *[]Task) {
	fmt.Printf("Task list:\t\n")
	for _, v := range *tasks {
		fmt.Printf("\tTitle: %s\n\tDescription: %s\n\tDueDate: %s\n\tCompleted: %t\n", v.Title, v.Description, v.DueDate, v.Completed)
	}
}

func main() {
	tasks := make([]Task, 0)

	AddTask(&tasks, Task{Title: "test-1", Description: "this is a test task", DueDate: "today", Completed: false})
	AddTask(&tasks, Task{Title: "test-2", Description: "this is another test task", DueDate: "tomorrow", Completed: false})

	ListTasks(&tasks)

	CompleteTask(&tasks, "test-1")
	ListTasks(&tasks)

	RemoveTask(&tasks, "test-2")
	ListTasks(&tasks)

	RemoveTask(&tasks, "test-2") // Attempt to remove again
}
