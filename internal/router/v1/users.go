package v1

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List Users"))
}
