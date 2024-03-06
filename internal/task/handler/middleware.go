package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sainath/todo-go-app/pkg/util"
)

func AuthenticationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationToken := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		userId, _ := strconv.Atoi(vars["userid"])
		if authorizationToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized access"))
			return
		}
		token := authorizationToken[len("Bearer "):]
		resp, err := http.Get("http://localhost:8081/username?id=" + strconv.Itoa(userId))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer resp.Body.Close()
		var userNameByte string
		json.NewDecoder(resp.Body).Decode(&userNameByte)

		if ok, err := util.ValidateToken(token); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		usernameToken := util.GetUsername(token)
		if usernameToken != userNameByte {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token not belongs to this user..!"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
