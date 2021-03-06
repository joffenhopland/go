package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"task_name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

// func allTasks() {
// 	task := Tasks{
// 		ID:         "1",
// 		TaskName:   "New projects",
// 		TaskDetail: "You must lead the project and finish it",
// 		Date:       "2022-01-22"}

// 	tasks = append(tasks, task)
// 	task1 := Tasks{
// 		ID:         "2",
// 		TaskName:   "Power projects",
// 		TaskDetail: "We need to hire more staff before the deadline",
// 		Date:       "2022-01-22"}

// 	tasks = append(tasks, task1)
// 	fmt.Println("Your tasks are", tasks)
// }

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am home page")
}
func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if params["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}
func createTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var task Tasks
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000))
	currentTime := time.Now().Format("01-02-2006")
	task.Date = currentTime
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(task)

}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	flag := false
	for index, item := range tasks {
		if item.ID == params["id"] { //0, 1,  2, 3, 4
			tasks = append(tasks[:index], tasks[index+1:]...)
			flag = true
			json.NewEncoder(w).Encode(map[string]string{"status": "Success"})
			fmt.Println("Deleted item", item.ID)

			return
		}
	}
	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
	}
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	flag := false
	for index, item := range tasks {
		if item.ID == params["id"] {
			fmt.Println("The id is", item.ID)
			tasks = append(tasks[:index], tasks[index+1:]...)
			var task Tasks
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			currentTime := time.Now().Format("01-02-2006")
			task.Date = currentTime
			tasks = append(tasks, task)
			flag = true
			json.NewEncoder(w).Encode(task)
			return

		}
	}

	if flag == false {
		json.NewEncoder(w).Encode(map[string]string{"status": "eror"})
	}

}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	// router.HandleFunc("/gettask/", getTask).Queries("id", "{id}").Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	// router.HandleFunc("/delete/", deleteTask).Queries("id", "{id}").Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	// allTasks()
	fmt.Println("Connection")
	handleRoutes()

}
