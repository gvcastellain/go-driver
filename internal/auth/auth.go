package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = "dsah9u33123123129g9sd9g9s9gsd9gsg0"

type Claims struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	jwt.RegisteredClaims
}

func createToken(authenticated Authenticated) (string, error) {
	expirationTime := time.Now().Add(30 * time.Minute)

	claims := &Claims{
		UserID:   authenticated.GetID(),
		UserName: authenticated.GetName(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

type Credentials struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (h *handler) auth(rw http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	u, err := h.authenticate(creds.UserName, creds.Password)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusUnauthorized)
	}

	token, err := createToken(u)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.Write([]byte(token))
}
