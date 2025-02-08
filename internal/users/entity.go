package users

import "time"

type User struct {
	ID         int64
	Name       string
	Login      string
	Password   string
	CreatedAt  time.Time
	ModifiedAt time.Time
	Deleted bool
	LastLogin time.Time
}
