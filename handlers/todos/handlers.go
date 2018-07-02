package todohandlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"todo/handlers"
	"todo/model/errors/input"
	"todo/model/task"
	"todo/services/todo"
	"todo/services/user"
)

type update struct {
	Success bool `json:"success"`
}

// TodoListHandler Handle a todo List
func TodoListHandler(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodGet && req.Method != http.MethodPost {
		http.NotFound(w, req)
		return
	}
	user := req.Header.Get("username")
	token := req.Header.Get("token")
	_, err := userservice.CheckSession(user, token)
	if err != nil {
		handlers.ErrorResponseHandler(w, req, err)
		return
	}
	encoder := json.NewEncoder(w)
	if req.Method == http.MethodGet {
		result, err := todoservice.GetList(user)
		if err != nil {
			handlers.ErrorResponseHandler(w, req, err)
		} else {
			w.Header().Set("content-type", "application/json")
			encoder.Encode(result)
		}
	} else {
		task := todotask.TodoTask{}
		json.NewDecoder(req.Body).Decode(&task)
		result, err := todoservice.AddTask(task.Desc, user)
		if err != nil {
			handlers.ErrorResponseHandler(w, req, err)
		} else {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(http.StatusCreated)
			encoder.Encode(result)
		}
	}
}

// TodoTaskHandler Handle a task in the todo list
func TodoTaskHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Method)
	if req.Method != http.MethodPut && req.Method != http.MethodDelete {
		http.NotFound(w, req)
		return
	}
	user := req.Header.Get("username")
	token := req.Header.Get("token")
	_, err := userservice.CheckSession(user, token)
	if err != nil {
		handlers.ErrorResponseHandler(w, req, err)
		return
	}
	ids := strings.Split(req.URL.Path, "/")[3]
	id, _ := strconv.ParseInt(ids, 10, 64)

	encoder := json.NewEncoder(w)
	if req.Method == http.MethodPut {
		var body interface{}
		json.NewDecoder(req.Body).Decode(&body)
		itemsMap := body.(map[string]interface{})
		done, found := itemsMap["done"].(bool)
		var result bool
		var err error
		if found {
			result, err = todoservice.UpdateTask(id, user, done)
		} else {
			err = inputerror.New([]string{"done"})
		}
		if err != nil {
			handlers.ErrorResponseHandler(w, req, err)
		} else {
			w.Header().Set("content-type", "application/json")
			//w.WriteHeader(http.StatusCreated)
			encoder.Encode(update{result})
		}
	} else {
		result, err := todoservice.DeleteTask(id, user)
		fmt.Println(result, err)
		if err != nil {
			handlers.ErrorResponseHandler(w, req, err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			//w.WriteHeader(http.StatusCreated)
			encoder.Encode(update{result})
		}
	}
}
