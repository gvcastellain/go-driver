package users

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	ErrPasswordRequired = errors.New("password cannot be blank")
	ErrNameRequired     = errors.New("name cannot be blank")
	ErrLoginRequired    = errors.New("login cannot be blank")
)

func New(name, login, password string) (*User, error) {
	u := User{Name: name, Login: login, ModifiedAt: time.Now()}

	err := u.SetPassword(password)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (u *User) Validade() error {
	if u.Name == "" {
		return ErrNameRequired
	}

	if u.Login == "" {
		return ErrLoginRequired
	}

	if u.Password == fmt.Sprintf("%x", md5.Sum([]byte(""))) {
		return ErrPasswordRequired
	}

	return nil
}

func encPass(p string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(p)))
}

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	u.Password = encPass(password)
	return nil
}

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}
