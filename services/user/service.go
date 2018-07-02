package userservice

import (
	"math/rand"
	"todo/model/errors/auth"
	"todo/model/user"
)

var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

const tokenLength int = 16

// Login validate USER loging
func Login(id string, password string) (string, error) {
	// lUser
	u := user.New(id)
	err := u.Load()
	var token string
	var logged bool
	if err != nil {
		token = ""
	} else {
		logged, err = u.CheckCredentials(password)
		if err != nil {
			return "", nil
		}
		if !logged {
			err = autherror.New()
		} else {
			token = genToken(tokenLength)
			err = u.SaveSesion(token)
		}
	}
	return token, err
}

// SignUp create a user in the system
func SignUp(id string, password string) (user.User, error) {

	u := user.New(id)
	u.Salt = genToken(16)
	u.EncryptPassword(password)

	u.Token = genToken(tokenLength)
	err := u.Save()
	return u, err
}

// Logout end user sesion
func Logout(id string, password string) (string, error) {
	// lUser
	u := user.New(id)
	err := u.Load()
	var token string
	var logged bool
	if err != nil {
		token = ""
	} else {
		logged, err = u.CheckCredentials(password)
		if err != nil {
			return "", err
		}
		if !logged {
			err = autherror.New()
		} else {
			err = u.SaveSesion("")
		}
	}
	return token, err
}

// CheckSesion Check if a user is loging in a sesion
func CheckSesion(id string, token string) (bool, error) {
	// lUser
	u := user.New(id)
	err := u.Load()
	var logged = false
	if err == nil {
		logged = token == u.Token
		err = nil
	}
	return logged, err

}

func createSesion(user *user.User) error {
	return nil
}

func genToken(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
