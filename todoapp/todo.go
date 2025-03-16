package main

import (
	"fmt"
	"net/http"
)

var item1 = "watch golang tutorial"
var item2 = "read a golang book"
var item3 = "build a project in go"
var actionItems = []string{item1, item2, item3}

func main() {
	//welcome messages
	fmt.Println("#####Welcome to our todo app #####")
	//handler
	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
	//call the function
	//printTasks()
	//printTasksWithParams(actionItems)
	//
	////add a new task
	//actionItems = addTask(actionItems, "learn go routines")
	//actionItems = addTask(actionItems, "learn channels")
	//printTasksWithParams(actionItems)
	/*
		fmt.Println(actionItems)
		//iterate over the slice and print each item
		//here we used _ to ignore the index -> similar to for each loop in java
		//you can still use the index if you want to -> when you need to access the indexes
		for index, task := range actionItems {
			//fmt.Println(index+1, ".", task) //this adds a space between the index and the task which is not good looking
			fmt.Printf("%d. %s\n", index+1, task)
		}

	*/
}

func showTasks(writer http.ResponseWriter, request *http.Request) {
	for _, item := range actionItems {
		fmt.Fprintln(writer, item)
	}
}

func helloUser(writer http.ResponseWriter, request *http.Request) {
	var greeting = "Hello, welcome to our todo app"
	fmt.Fprintf(writer, greeting)

}

func printTasks() {
	fmt.Println("###### Tasks ######")
	for index, task := range actionItems {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

// parameterized function
func printTasksWithParams(tasks []string) {
	fmt.Println("###### Tasks ######")
	for index, task := range tasks {
		fmt.Printf("%d. %s\n", index+1, task)
	}
}

// lists in go
func addTask(tasks []string, newTask string) []string {
	var updatedTaskItems = append(tasks, newTask)
	return updatedTaskItems
}
