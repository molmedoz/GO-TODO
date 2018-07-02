package userhandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todo/handlers"
	"todo/model/errors/auth"
	"todo/model/errors/input"
	"todo/services/user"
)

type loginForm struct {
	user     string
	password string
}
type signUpForm struct {
	user      string
	password  string
	password2 string
}

type authResponse struct {
	User  string `json:"user"`
	Token string `json:"token"`
}

// LoginHandler Handler for create the loggin
func LoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	var body interface{}
	json.NewDecoder(req.Body).Decode(&body)
	user, password, err := validateLoginBody(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err)
	}
	form := loginForm{}
	json.NewDecoder(req.Body).Decode(&form)
	token, err := userservice.Login(user, password)
	if err != nil {
		switch err.(type) {
		case *autherror.AuthError:
			{
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		default:
			{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	encoder.Encode(authResponse{User: user, Token: token})

}

// SignUpHandler Handle a request for creating a new user
func SignUpHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	var body interface{}
	json.NewDecoder(req.Body).Decode(&body)
	user, password, err := validateSignUpBody(body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		encoder.Encode(err)
	}
	form := loginForm{}
	json.NewDecoder(req.Body).Decode(&form)
	u, err := userservice.SignUp(user, password)
	if err != nil {
		switch err.(type) {
		case *autherror.AuthError:
			{
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		default:
			{
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
	encoder.Encode(authResponse{User: user, Token: u.Token})
}

// LogoutHandler Handle a request for creating a new user
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.NotFound(w, req)
		return
	}
	id := req.Header.Get("username")
	token := req.Header.Get("token")
	_, err := userservice.Logout(id, token)
	if err != nil {
		handlers.ErrorResponseHandler(w, req, err)
	}

}
func validateLoginBody(body interface{}) (string, string, error) {
	if body == nil {
		return "", "", inputerror.New([]string{"user", "password"})
	}
	itemsMap := body.(map[string]interface{})
	fmt.Println(body, itemsMap)
	fields := []string{}
	user, found := itemsMap["user"].(string)
	if !found || user == "" {
		fields = append(fields, "user")
	}
	password, found := itemsMap["password"].(string)
	if !found || password == "" {
		fields = append(fields, "password")
	}
	if len(fields) > 0 {
		return "", "", inputerror.New(fields)
	}
	return user, password, nil
}

func validateSignUpBody(body interface{}) (string, string, error) {
	if body == nil {
		return "", "", inputerror.New([]string{"user", "password", "password2"})
	}
	itemsMap := body.(map[string]interface{})
	fields := []string{}
	user, found := itemsMap["user"].(string)
	if !found || user == "" {
		fields = append(fields, "user")
	}
	password, found := itemsMap["password"].(string)
	if !found || password == "" {
		fields = append(fields, "password")
	}
	password2, found := itemsMap["password2"].(string)
	if !found || password == "" {
		fields = append(fields, "password2")
	}
	if len(fields) > 0 {
		return "", "", inputerror.New(fields)
	}
	if password != password2 {
		fields = []string{"password", "password2"}
		return "", "", inputerror.New(fields)
	}
	return user, password, nil
}
