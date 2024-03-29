package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sainath/todo-go-app/internal/user/model"
	"github.com/sainath/todo-go-app/pkg/util"
)

func AuthenticationMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationToken := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		userId, _ := strconv.Atoi(vars["id"])
		if authorizationToken == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("User Unauthorized"))
			return
		}
		token := authorizationToken[len("Bearer "):]
		var user model.User
		db.First(&user, userId)

		if ok, err := util.ValidateToken(token); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		usernameToken := util.GetUsername(token)
		if usernameToken != user.Email {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token not belongs to this user..!"))
			return
		}

		handler.ServeHTTP(w, r)
	})
}
