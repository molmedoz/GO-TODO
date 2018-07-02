package user

import (
	"todo/model"

	"golang.org/x/crypto/scrypt"
)

// User {
// 	ID       string User id, represent the alias
// 	Salt     string Public key for the password
// 	Password string Encrypted password
// }
type User struct {
	ID       string `json:"id"`
	Salt     string `json:"-"`
	Password string `json:"-"`
	Token    string `json:"token,omitempty"`
}

const (
	pwSaltBytes = 32
	pwHashBytes = 64
)

func generateHashSalt(password string, salt string) (string, error) {
	hash, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 32)
	return string(hash), err

}

// ValidatePassword Check if a given password is valid
func validatePassword(salt string, hash string, password string) (bool, error) {
	hashed, err := generateHashSalt(password, salt)
	if err != nil {
		return false, err
	}
	return hashed == hash, err
}

// New Generate a new user
func New(id string) User {
	return User{ID: id}
}

// EncryptPassword Create salt, hash for Authentication
func (u *User) EncryptPassword(password string) error {
	hash, err := generateHashSalt(password, u.Salt)
	if err == nil {
		u.Password = hash
	}
	return err
}

// CheckCredentials Check the credentials of a user
func (u *User) CheckCredentials(secret string) (bool, error) {
	if u.ID == "" {
		return false, nil
	}
	// logger, err := validatePassword(u.Salt, u.Password, secret)
	return validatePassword(u.Salt, u.Password, secret)
}

// Load a user from the database
func (u *User) Load() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}
	rows, err := db.Query("SELECT salt, password, token FROM USERS WHERE id=?", u.ID)
	if err != nil {
		return err
	}
	defer rows.Close()
	n := 0
	for rows.Next() {
		n++
		rows.Scan(&u.Salt, &u.Password, &u.Token)
		break
	}
	if n != 1 {
		u.ID = ""
		u.Password = ""
		u.Salt = ""
		u.Token = ""
	}

	return nil
}

// Save user in de database
func (u *User) Save() error {
	db, _ := database.GetConnection()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stm, err := tx.Prepare("INSERT INTO USERS (id, salt, password, token) VALUES (?, ?, ?, ?)")
	if err != nil {
		return nil
	}
	_, err = stm.Exec(u.ID, u.Salt, u.Password, u.Token)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

// SaveSesion Save save the user in the database
func (u *User) SaveSesion(token string) error {
	db, _ := database.GetConnection()
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	stm, err := tx.Prepare("UPDATE USERS SET token=? WHERE id=?")
	if err != nil {
		return nil
	}
	_, err = stm.Exec(token, u.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	u.Token = token
	return tx.Commit()
}
