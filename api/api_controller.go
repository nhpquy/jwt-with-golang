package main

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

const JwtTokenName = "jwt_token"
const LoginHost = "http://localhost:3000"

type ApiController struct{}

func (apiController ApiController) Index(w http.ResponseWriter, r *http.Request) {
	tokenString, err := getJwt(r)

	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, LoginHost, http.StatusSeeOther)
		return
	}

	if tokenString != "" {
		fmt.Println(tokenString)

		token, err := Sign(tokenString)

		if err != nil {
			fmt.Println(err)
			return
		}

		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			cookie := http.Cookie{
				Name:  JwtTokenName,
				Value: tokenString,
			}
			http.SetCookie(w, &cookie)
			w.WriteHeader(http.StatusOK)
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
}

func getJwt(r *http.Request) (tokenString string, err error) {
	tokenString = r.URL.Query().Get("authentication")

	if tokenString == "" {
		cookie, err := r.Cookie(JwtTokenName)

		if err != nil {
			return "", err
		}

		tokenString = cookie.Value
	} else {
		err = nil
	}
	return tokenString, err
}
