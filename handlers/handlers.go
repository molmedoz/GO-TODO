package handlers

import (
	"encoding/json"
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
		encoder := json.NewEncoder(w)
		w.Header().Set("content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err)
		return
	case *notfounderror.NotFoundError:
		w.WriteHeader(http.StatusNotFound)
		return
	case *autherror.AuthError:
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
}
