package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/handlers/todos"
	"todo/handlers/users"
	"todo/model"

	_ "github.com/mattn/go-sqlite3"
)

/**
 * HealthCheck structure
 * @type {[type]}
 */
type healthCheck struct {
	DBStatus bool
}

func main() {

	mux := http.NewServeMux()

	db, err := database.GetConnection()
	fmt.Println(db, err)
	mux.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("request")
		connected := db.Ping() != nil
		encoder := json.NewEncoder(w)
		data := healthCheck{connected}
		// data := &TodoTask{1, "Store in DB", false, "Manu"}
		//
		encoder.Encode(data)
	})
	mux.HandleFunc("/login", userhandler.LoginHandler)
	mux.HandleFunc("/signup", userhandler.SignUpHandler)
	mux.HandleFunc("/logout", userhandler.Logout)
	mux.HandleFunc("/api/todo", todohandlers.TodoListHandler) //func(w http.ResponseWriter, req *http.Request) {
	mux.HandleFunc("/api/todo/", todohandlers.TodoTaskHandler)

	//})
	// mux.H

	fmt.Println(http.ListenAndServe(":8080", mux))
}
