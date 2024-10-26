package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/VsenseTechnologies/skf_plc_http_server/presentation/model/response"
	"github.com/golang-jwt/jwt/v5"
)

func AdminAuthenticationMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		jwtSecreteKey := os.Getenv("SECRETE_KEY")

		if jwtSecreteKey == "" {
			log.Printf("missing env variable SECRETE_KEY")
			response := response.StatusMessage{
				Message: "missing enviroment variable",
			}

			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		cookie, err := r.Cookie("token")

		requestToken := strings.Split(cookie.String(), "=")[1]

		if err != nil {
			if err == http.ErrNoCookie {
				response := response.StatusMessage{
					Message: "token not found",
				}
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response)
				return
			}
			log.Printf("error occurred while getting the token from cookie, host -> %s", r.Host)
			response := response.StatusMessage{
				Message: "error occurred while getting token from cookie",
			}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(response)
			return
		}

		_, error := jwt.Parse(requestToken, func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodHS256 {
				return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
			}

			return jwtSecreteKey, nil
		})

		fmt.Println(error.Error())

		if error != nil {
			response := response.StatusMessage{
				Message: "invalid token",
			}

			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
