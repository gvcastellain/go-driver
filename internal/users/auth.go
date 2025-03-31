package users

import "time"

func (h *handler) autenticate(login, password string) (*User, error) {
	stmt := `select * from "users" where login = $1 and password = $2`

	row := h.db.QueryRow(stmt, login, encPass(password))

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.CreatedAt, &u.ModifiedAt, &u.Deleted, &u.LastLogin)

	if err != nil {
		return nil, err
	}

	return &u, nil

}

func (h *handler) updateLastLogin(u *User) error {
	u.LastLogin = time.Now()

	return Update(h.db, u.ID, u)
}

func Autenticate(login, password string) (u *User, err error) {
	u, err = gloabalHandler.autenticate(login, password)
	if err != nil {
		return
	}

	err = gloabalHandler.updateLastLogin(u)
	if err != nil {
		return nil, err
	}

	return
}
