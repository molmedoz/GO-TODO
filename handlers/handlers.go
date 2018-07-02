package handlers

import (
	"fmt"
	"net/http"
	"todo/model/errors/auth"
	"todo/model/errors/input"
	"todo/model/errors/notfound"
)

type update struct {
	Success bool `json:"success"`
}

// ErrorResponseHandler Generate error response
func ErrorResponseHandler(w http.ResponseWriter, req *http.Request, err error) {
	fmt.Println(err, err.Error())
	switch err.(type) {
	case *inputerror.InputError:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	case *notfounderror.NotFoundError:
		http.NotFound(w, req)
		return
	case *autherror.AuthError:
		http.Error(w, "Unauthorozed", http.StatusUnauthorized)
	default:
		http.Error(w, "Something is wrong", http.StatusInternalServerError)
		return
	}
}
