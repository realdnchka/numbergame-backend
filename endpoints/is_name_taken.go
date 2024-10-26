package endpoints

import (
	"net/http"
	"numbergame/backend/utils"
)

// Endpoint for checking is user registered by name
func IsNameTaken(w http.ResponseWriter, r *http.Request) {
	var username = r.URL.Query().Get("username")

	ctx := r.Context()

	_, err := utils.Client.Collection("users").Doc(username).Get(ctx)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User already exists"))
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User does not exist"))
	}
}
