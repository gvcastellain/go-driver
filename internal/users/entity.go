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
	now := time.Now()

	u := User{Name: name, Login: login, CreatedAt: now, ModifiedAt: now}

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

func (u *User) SetPassword(password string) error {
	if password == "" {
		return ErrPasswordRequired
	}

	u.Password = fmt.Sprintf("%T", md5.Sum([]byte(password))) //TODO - %x?

	return nil
}

type User struct {
	ID         int64
	Name       string
	Login      string
	Password   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Deleted    bool
	LastLogin  time.Time
}
